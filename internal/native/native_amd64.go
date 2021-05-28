/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package native

import (
    `unsafe`
)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __i64toa(out *byte, val int64) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __u64toa(out *byte, val uint64) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __f64toa(out *byte, val float64) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __lzero(p unsafe.Pointer, n int) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __lquote(buf *string, off int) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __lspace(sp unsafe.Pointer, nb int, off int) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __value(s unsafe.Pointer, n int, p int, v *JsonState) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __vstring(s *string, p *int, v *JsonState)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __vnumber(s *string, p *int, v *JsonState)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __vsigned(s *string, p *int, v *JsonState)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __vunsigned(s *string, p *int, v *JsonState)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __skip_one(s *string, p *int, m *StateMachine) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __skip_array(s *string, p *int, m *StateMachine) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __skip_object(s *string, p *int, m *StateMachine) (ret int)

//go:nosplit
//go:noescape
//goland:noinspection GoUnusedParameter
func __unquote(s unsafe.Pointer, nb int, dp unsafe.Pointer, ep *int, flags uint64) (ret int)