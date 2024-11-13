// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.uber.org/multierr"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

func (s *Server) decodeCreateUserRequest(r *http.Request) (
	req CreateUserReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	req = &CreateUserReqEmptyBody{}
	if _, ok := r.Header["Content-Type"]; !ok && r.ContentLength == 0 {
		return req, close, nil
	}
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, nil
		}

		d := jx.DecodeBytes(buf)

		var request CreateUserApplicationJSON
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, close, err
		}
		return &request, close, nil
	case ct == "application/x-www-form-urlencoded":
		if r.ContentLength == 0 {
			return req, close, nil
		}
		form, err := ht.ParseForm(r)
		if err != nil {
			return req, close, errors.Wrap(err, "parse form")
		}

		var request CreateUserApplicationXWwwFormUrlencoded
		{
			var unwrapped User
			q := uri.NewQueryDecoder(form)
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "id",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						var unwrappedDotIDVal int64
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt64(val)
							if err != nil {
								return err
							}

							unwrappedDotIDVal = c
							return nil
						}(); err != nil {
							return err
						}
						unwrapped.ID.SetTo(unwrappedDotIDVal)
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"id\"")
					}
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "nickname",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						var unwrappedDotNicknameVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							unwrappedDotNicknameVal = c
							return nil
						}(); err != nil {
							return err
						}
						unwrapped.Nickname.SetTo(unwrappedDotNicknameVal)
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"nickname\"")
					}
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "name",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						var unwrappedDotNameVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							unwrappedDotNameVal = c
							return nil
						}(); err != nil {
							return err
						}
						unwrapped.Name.SetTo(unwrappedDotNameVal)
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"name\"")
					}
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "surname",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						var unwrappedDotSurnameVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							unwrappedDotSurnameVal = c
							return nil
						}(); err != nil {
							return err
						}
						unwrapped.Surname.SetTo(unwrappedDotSurnameVal)
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"surname\"")
					}
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "hashedPass",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						var unwrappedDotHashedPassVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							unwrappedDotHashedPassVal = c
							return nil
						}(); err != nil {
							return err
						}
						unwrapped.HashedPass.SetTo(unwrappedDotHashedPassVal)
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"hashedPass\"")
					}
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "userStatus",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						var unwrappedDotUserStatusVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							unwrappedDotUserStatusVal = c
							return nil
						}(); err != nil {
							return err
						}
						unwrapped.UserStatus.SetTo(unwrappedDotUserStatusVal)
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"userStatus\"")
					}
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "date",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						var unwrappedDotDateVal time.Time
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToDate(val)
							if err != nil {
								return err
							}

							unwrappedDotDateVal = c
							return nil
						}(); err != nil {
							return err
						}
						unwrapped.Date.SetTo(unwrappedDotDateVal)
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"date\"")
					}
				}
			}
			request = CreateUserApplicationXWwwFormUrlencoded(unwrapped)
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeCreateUsersWithListInputRequest(r *http.Request) (
	req []User,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	if _, ok := r.Header["Content-Type"]; !ok && r.ContentLength == 0 {
		return req, close, nil
	}
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, nil
		}

		d := jx.DecodeBytes(buf)

		var request []User
		if err := func() error {
			request = make([]User, 0)
			if err := d.Arr(func(d *jx.Decoder) error {
				var elem User
				if err := elem.Decode(d); err != nil {
					return err
				}
				request = append(request, elem)
				return nil
			}); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, close, err
		}
		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeUpdateUserRequest(r *http.Request) (
	req UpdateUserReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	req = &UpdateUserReqEmptyBody{}
	if _, ok := r.Header["Content-Type"]; !ok && r.ContentLength == 0 {
		return req, close, nil
	}
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, nil
		}

		d := jx.DecodeBytes(buf)

		var request UpdateUserApplicationJSON
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			if err := d.Skip(); err != io.EOF {
				return errors.New("unexpected trailing data")
			}
			return nil
		}(); err != nil {
			err = &ogenerrors.DecodeBodyError{
				ContentType: ct,
				Body:        buf,
				Err:         err,
			}
			return req, close, err
		}
		return &request, close, nil
	case ct == "application/x-www-form-urlencoded":
		if r.ContentLength == 0 {
			return req, close, nil
		}
		form, err := ht.ParseForm(r)
		if err != nil {
			return req, close, errors.Wrap(err, "parse form")
		}

		var request UpdateUserApplicationXWwwFormUrlencoded
		{
			var unwrapped User
			q := uri.NewQueryDecoder(form)
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "id",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						var unwrappedDotIDVal int64
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt64(val)
							if err != nil {
								return err
							}

							unwrappedDotIDVal = c
							return nil
						}(); err != nil {
							return err
						}
						unwrapped.ID.SetTo(unwrappedDotIDVal)
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"id\"")
					}
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "nickname",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						var unwrappedDotNicknameVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							unwrappedDotNicknameVal = c
							return nil
						}(); err != nil {
							return err
						}
						unwrapped.Nickname.SetTo(unwrappedDotNicknameVal)
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"nickname\"")
					}
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "name",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						var unwrappedDotNameVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							unwrappedDotNameVal = c
							return nil
						}(); err != nil {
							return err
						}
						unwrapped.Name.SetTo(unwrappedDotNameVal)
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"name\"")
					}
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "surname",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						var unwrappedDotSurnameVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							unwrappedDotSurnameVal = c
							return nil
						}(); err != nil {
							return err
						}
						unwrapped.Surname.SetTo(unwrappedDotSurnameVal)
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"surname\"")
					}
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "hashedPass",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						var unwrappedDotHashedPassVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							unwrappedDotHashedPassVal = c
							return nil
						}(); err != nil {
							return err
						}
						unwrapped.HashedPass.SetTo(unwrappedDotHashedPassVal)
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"hashedPass\"")
					}
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "userStatus",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						var unwrappedDotUserStatusVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							unwrappedDotUserStatusVal = c
							return nil
						}(); err != nil {
							return err
						}
						unwrapped.UserStatus.SetTo(unwrappedDotUserStatusVal)
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"userStatus\"")
					}
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "date",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						var unwrappedDotDateVal time.Time
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToDate(val)
							if err != nil {
								return err
							}

							unwrappedDotDateVal = c
							return nil
						}(); err != nil {
							return err
						}
						unwrapped.Date.SetTo(unwrappedDotDateVal)
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"date\"")
					}
				}
			}
			request = UpdateUserApplicationXWwwFormUrlencoded(unwrapped)
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}