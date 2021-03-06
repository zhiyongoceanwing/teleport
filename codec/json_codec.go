// Copyright 2015-2017 HenryLee. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codec

import (
	"bytes"
	"encoding/json"
	"io"
)

func init() {
	Reg(new(JsonCodec))
}

// JsonCodec json codec
type JsonCodec struct{}

// Name returns codec name
func (j *JsonCodec) Name() string {
	return "json"
}

//
// Id returns codec id
func (j *JsonCodec) Id() byte {
	return 'j'
}

// JsonEncoder json decoder
type JsonEncoder struct {
	writer io.Writer
}

// NewEncoder returns a new json encoder that writes to writer.
func (*JsonCodec) NewEncoder(writer io.Writer) Encoder {
	return &JsonEncoder{writer}
}

// Encode writes the json encoding of v to the writer.
func (p *JsonEncoder) Encode(v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = p.writer.Write(b)
	return err
}

// JsonDecoder json decoder
type JsonDecoder struct {
	limitReader io.Reader
	buf         *bytes.Buffer
}

// NewDecoder returns a new json decoder that reads from limit reader.
func (*JsonCodec) NewDecoder(limitReader io.Reader) Decoder {
	return &JsonDecoder{
		limitReader: limitReader,
		buf:         bytes.NewBuffer(make([]byte, 0, bytes.MinRead)),
	}
}

// Decode reads the next json-encoded value from its
// input and stores it in the value pointed to by v.
func (p *JsonDecoder) Decode(v interface{}) error {
	p.buf.Reset()
	_, err := p.buf.ReadFrom(p.limitReader)
	if err != nil && err != io.ErrUnexpectedEOF {
		return err
	}
	return json.Unmarshal(p.buf.Bytes(), v)
}
