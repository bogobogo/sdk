package protocol

import (
	"io"

	"github.com/bblfsh/sdk/protocol/jsonlines"
	"github.com/bblfsh/sdk/uast"
	"github.com/bblfsh/sdk/uast/ann"

	"srcd.works/go-errors.v0"
)

var (
	ErrDecodingRequest  = errors.NewKind("error decoding request")
	ErrEncodingResponse = errors.NewKind("error encoding response")

	noErrClose = errors.NewKind("clean close").New()
)

type Server struct {
	In       io.Reader
	Out      io.Writer
	Native   *NativeClient
	ToNoder  uast.ToNoder
	Annotate *ann.Rule

	err  error
	done chan struct{}
}

func (s *Server) Start() error {
	s.done = make(chan struct{})
	go func() {
		s.serve()
	}()

	return nil
}

func (s *Server) Wait() error {
	<-s.done
	return s.err
}

func (s *Server) serve() {
	enc := jsonlines.NewEncoder(s.Out)
	dec := jsonlines.NewDecoder(s.In)
	for {
		if err := s.process(enc, dec); err != nil {
			if err == noErrClose {
				break
			}

			s.err = err
			break
		}
	}

	close(s.done)
}

func (s *Server) process(enc jsonlines.Encoder, dec jsonlines.Decoder) error {
	req := &ParseUASTRequest{}
	if err := dec.Decode(req); err != nil {
		if err == io.EOF {
			return noErrClose
		}

		return s.error(enc, ErrDecodingRequest.Wrap(err))
	}

	resp, err := s.Native.ParseNativeAST(&ParseNativeASTRequest{
		Content: req.Content,
	})
	if err != nil {
		return s.error(enc, err)
	}

	var node *uast.Node
	if resp.Status != Fatal {
		node, err = s.ToNoder.ToNode(resp.AST)
		if err != nil {
			return s.error(enc, err)
		}

		if err := s.Annotate.Apply(node); err != nil {
			return s.error(enc, err)
		}
	}

	uastResp := &ParseUASTResponse{
		Status: resp.Status,
		Errors: resp.Errors,
		UAST:   node,
	}
	if err := enc.Encode(uastResp); err != nil {
		return ErrEncodingResponse.Wrap(err)
	}

	return nil
}

func (s *Server) error(enc jsonlines.Encoder, err error) error {
	if err := enc.Encode(newFatalResponse(err)); err != nil {
		return ErrEncodingResponse.Wrap(err)
	}

	return nil
}

func newFatalResponse(err error) *ParseUASTResponse {
	return &ParseUASTResponse{
		Status: Fatal,
		Errors: []string{err.Error()},
	}
}
