// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson1fe40eefDecodeWumeComposerInternalPkgModels(in *jlexer.Lexer, out *UsersDataAnswer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "status":
			out.Status = int(in.Int())
		case "message":
			out.Message = string(in.String())
		case "data":
			(out.Data).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1fe40eefEncodeWumeComposerInternalPkgModels(out *jwriter.Writer, in UsersDataAnswer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Status))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		(in.Data).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UsersDataAnswer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1fe40eefEncodeWumeComposerInternalPkgModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UsersDataAnswer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1fe40eefEncodeWumeComposerInternalPkgModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UsersDataAnswer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1fe40eefDecodeWumeComposerInternalPkgModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UsersDataAnswer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1fe40eefDecodeWumeComposerInternalPkgModels(l, v)
}
func easyjson1fe40eefDecodeWumeComposerInternalPkgModels1(in *jlexer.Lexer, out *UserDataAnswer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "status":
			out.Status = int(in.Int())
		case "message":
			out.Message = string(in.String())
		case "data":
			(out.Data).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1fe40eefEncodeWumeComposerInternalPkgModels1(out *jwriter.Writer, in UserDataAnswer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Status))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		(in.Data).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserDataAnswer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1fe40eefEncodeWumeComposerInternalPkgModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserDataAnswer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1fe40eefEncodeWumeComposerInternalPkgModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserDataAnswer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1fe40eefDecodeWumeComposerInternalPkgModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserDataAnswer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1fe40eefDecodeWumeComposerInternalPkgModels1(l, v)
}
func easyjson1fe40eefDecodeWumeComposerInternalPkgModels2(in *jlexer.Lexer, out *UploadAvatarAnswer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "status":
			out.Status = int(in.Int())
		case "message":
			out.Message = string(in.String())
		case "data":
			out.Data = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1fe40eefEncodeWumeComposerInternalPkgModels2(out *jwriter.Writer, in UploadAvatarAnswer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Status))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		out.String(string(in.Data))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UploadAvatarAnswer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1fe40eefEncodeWumeComposerInternalPkgModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UploadAvatarAnswer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1fe40eefEncodeWumeComposerInternalPkgModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UploadAvatarAnswer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1fe40eefDecodeWumeComposerInternalPkgModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UploadAvatarAnswer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1fe40eefDecodeWumeComposerInternalPkgModels2(l, v)
}
func easyjson1fe40eefDecodeWumeComposerInternalPkgModels3(in *jlexer.Lexer, out *MessageAnswer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "status":
			out.Status = int(in.Int())
		case "message":
			out.Message = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1fe40eefEncodeWumeComposerInternalPkgModels3(out *jwriter.Writer, in MessageAnswer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Status))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MessageAnswer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1fe40eefEncodeWumeComposerInternalPkgModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MessageAnswer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1fe40eefEncodeWumeComposerInternalPkgModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MessageAnswer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1fe40eefDecodeWumeComposerInternalPkgModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MessageAnswer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1fe40eefDecodeWumeComposerInternalPkgModels3(l, v)
}
func easyjson1fe40eefDecodeWumeComposerInternalPkgModels4(in *jlexer.Lexer, out *IncorrectDataAnswer) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "status":
			out.Status = int(in.Int())
		case "message":
			out.Message = string(in.String())
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				in.Delim('[')
				if out.Data == nil {
					if !in.IsDelim(']') {
						out.Data = make([]string, 0, 4)
					} else {
						out.Data = []string{}
					}
				} else {
					out.Data = (out.Data)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Data = append(out.Data, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1fe40eefEncodeWumeComposerInternalPkgModels4(out *jwriter.Writer, in IncorrectDataAnswer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Status))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		if in.Data == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Data {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v IncorrectDataAnswer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1fe40eefEncodeWumeComposerInternalPkgModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IncorrectDataAnswer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1fe40eefEncodeWumeComposerInternalPkgModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *IncorrectDataAnswer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1fe40eefDecodeWumeComposerInternalPkgModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IncorrectDataAnswer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1fe40eefDecodeWumeComposerInternalPkgModels4(l, v)
}
