// Code generated by 'ccgo -D__environ=environ -export-externs X -hide __syscall0,__syscall1,__syscall2,__syscall3,__syscall4,__syscall5,__syscall6 -nostdinc -nostdlib -o ../musl_windows_386.go -pkgname libc -Iarch/i386 -Iarch/generic -Iobj/src/internal -Isrc/include -Isrc/internal -Iobj/include -Iinclude copyright.c src/ctype/isalnum.c src/ctype/isalpha.c src/ctype/isdigit.c src/ctype/islower.c src/ctype/isprint.c src/ctype/isspace.c src/ctype/isxdigit.c src/env/putenv.c src/env/setenv.c src/env/unsetenv.c src/string/strchrnul.c', DO NOT EDIT.

package libc

import (
	"math"
	"reflect"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ unsafe.Pointer

// musl as a whole is licensed under the following standard MIT license:
//
// ----------------------------------------------------------------------
// Copyright © 2005-2020 Rich Felker, et al.
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
// CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
// SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
// ----------------------------------------------------------------------
//
// Authors/contributors include:
//
// A. Wilcox
// Ada Worcester
// Alex Dowad
// Alex Suykov
// Alexander Monakov
// Andre McCurdy
// Andrew Kelley
// Anthony G. Basile
// Aric Belsito
// Arvid Picciani
// Bartosz Brachaczek
// Benjamin Peterson
// Bobby Bingham
// Boris Brezillon
// Brent Cook
// Chris Spiegel
// Clément Vasseur
// Daniel Micay
// Daniel Sabogal
// Daurnimator
// David Carlier
// David Edelsohn
// Denys Vlasenko
// Dmitry Ivanov
// Dmitry V. Levin
// Drew DeVault
// Emil Renner Berthing
// Fangrui Song
// Felix Fietkau
// Felix Janda
// Gianluca Anzolin
// Hauke Mehrtens
// He X
// Hiltjo Posthuma
// Isaac Dunham
// Jaydeep Patil
// Jens Gustedt
// Jeremy Huntwork
// Jo-Philipp Wich
// Joakim Sindholt
// John Spencer
// Julien Ramseier
// Justin Cormack
// Kaarle Ritvanen
// Khem Raj
// Kylie McClain
// Leah Neukirchen
// Luca Barbato
// Luka Perkov
// M Farkas-Dyck (Strake)
// Mahesh Bodapati
// Markus Wichmann
// Masanori Ogino
// Michael Clark
// Michael Forney
// Mikhail Kremnyov
// Natanael Copa
// Nicholas J. Kain
// orc
// Pascal Cuoq
// Patrick Oppenlander
// Petr Hosek
// Petr Skocik
// Pierre Carrier
// Reini Urban
// Rich Felker
// Richard Pennington
// Ryan Fairfax
// Samuel Holland
// Segev Finer
// Shiz
// sin
// Solar Designer
// Stefan Kristiansson
// Stefan O'Rear
// Szabolcs Nagy
// Timo Teräs
// Trutz Behn
// Valentin Ochs
// Will Dietz
// William Haddon
// William Pitcock
//
// Portions of this software are derived from third-party works licensed
// under terms compatible with the above MIT license:
//
// The TRE regular expression implementation (src/regex/reg* and
// src/regex/tre*) is Copyright © 2001-2008 Ville Laurikari and licensed
// under a 2-clause BSD license (license text in the source files). The
// included version has been heavily modified by Rich Felker in 2012, in
// the interests of size, simplicity, and namespace cleanliness.
//
// Much of the math library code (src/math/* and src/complex/*) is
// Copyright © 1993,2004 Sun Microsystems or
// Copyright © 2003-2011 David Schultz or
// Copyright © 2003-2009 Steven G. Kargl or
// Copyright © 2003-2009 Bruce D. Evans or
// Copyright © 2008 Stephen L. Moshier or
// Copyright © 2017-2018 Arm Limited
// and labelled as such in comments in the individual source files. All
// have been licensed under extremely permissive terms.
//
// The ARM memcpy code (src/string/arm/memcpy.S) is Copyright © 2008
// The Android Open Source Project and is licensed under a two-clause BSD
// license. It was taken from Bionic libc, used on Android.
//
// The AArch64 memcpy and memset code (src/string/aarch64/*) are
// Copyright © 1999-2019, Arm Limited.
//
// The implementation of DES for crypt (src/crypt/crypt_des.c) is
// Copyright © 1994 David Burren. It is licensed under a BSD license.
//
// The implementation of blowfish crypt (src/crypt/crypt_blowfish.c) was
// originally written by Solar Designer and placed into the public
// domain. The code also comes with a fallback permissive license for use
// in jurisdictions that may not recognize the public domain.
//
// The smoothsort implementation (src/stdlib/qsort.c) is Copyright © 2011
// Valentin Ochs and is licensed under an MIT-style license.
//
// The x86_64 port was written by Nicholas J. Kain and is licensed under
// the standard MIT terms.
//
// The mips and microblaze ports were originally written by Richard
// Pennington for use in the ellcc project. The original code was adapted
// by Rich Felker for build system and code conventions during upstream
// integration. It is licensed under the standard MIT terms.
//
// The mips64 port was contributed by Imagination Technologies and is
// licensed under the standard MIT terms.
//
// The powerpc port was also originally written by Richard Pennington,
// and later supplemented and integrated by John Spencer. It is licensed
// under the standard MIT terms.
//
// All other files which have no copyright comments are original works
// produced specifically for use as part of this library, written either
// by Rich Felker, the main author of the library, or by one or more
// contibutors listed above. Details on authorship of individual files
// can be found in the git version control history of the project. The
// omission of copyright and license comments in each file is in the
// interest of source tree size.
//
// In addition, permission is hereby granted for all public header files
// (include/* and arch/*/bits/*) and crt files intended to be linked into
// applications (crt/*, ldso/dlstart.c, and arch/*/crt_arch.h) to omit
// the copyright notice and permission notice otherwise required by the
// license, and to use these files without any requirement of
// attribution. These files include substantial contributions from:
//
// Bobby Bingham
// John Spencer
// Nicholas J. Kain
// Rich Felker
// Richard Pennington
// Stefan Kristiansson
// Szabolcs Nagy
//
// all of whom have explicitly granted such permission.
//
// This file previously contained text expressing a belief that most of
// the files covered by the above exception were sufficiently trivial not
// to be subject to copyright, resulting in confusion over whether it
// negated the permissions granted in the license. In the spirit of
// permissive licensing, and of not having licensing issues being an
// obstacle to adoption, that text has been removed.
const ( /* copyright.c:194:1: */
	__musl__copyright__ = 0
)

type ptrdiff_t = int32 /* <builtin>:3:26 */

type size_t = uint32 /* <builtin>:9:23 */

type wchar_t = uint16 /* <builtin>:15:24 */

type va_list = uintptr /* <builtin>:51:27 */

type locale_t = uintptr /* alltypes.h:366:32 */

func Xisalnum(tls *TLS, c int32) int32 { /* isalnum.c:3:5: */
	return (Bool32((func() int32 {
		if 0 != 0 {
			return Xisalpha(tls, c)
		}
		return (Bool32((((uint32(c)) | uint32(32)) - uint32('a')) < uint32(26)))
	}() != 0) || (func() int32 {
		if 0 != 0 {
			return Xisdigit(tls, c)
		}
		return (Bool32(((uint32(c)) - uint32('0')) < uint32(10)))
	}() != 0)))
}

func X__isalnum_l(tls *TLS, c int32, l locale_t) int32 { /* isalnum.c:8:5: */
	return Xisalnum(tls, c)
}

func Xisalpha(tls *TLS, c int32) int32 { /* isalpha.c:4:5: */
	return (Bool32(((uint32(c) | uint32(32)) - uint32('a')) < uint32(26)))
}

func X__isalpha_l(tls *TLS, c int32, l locale_t) int32 { /* isalpha.c:9:5: */
	return Xisalpha(tls, c)
}

func Xisdigit(tls *TLS, c int32) int32 { /* isdigit.c:4:5: */
	return (Bool32((uint32(c) - uint32('0')) < uint32(10)))
}

func X__isdigit_l(tls *TLS, c int32, l locale_t) int32 { /* isdigit.c:9:5: */
	return Xisdigit(tls, c)
}

func Xislower(tls *TLS, c int32) int32 { /* islower.c:4:5: */
	return (Bool32((uint32(c) - uint32('a')) < uint32(26)))
}

func X__islower_l(tls *TLS, c int32, l locale_t) int32 { /* islower.c:9:5: */
	return Xislower(tls, c)
}

func Xisprint(tls *TLS, c int32) int32 { /* isprint.c:4:5: */
	return (Bool32((uint32(c) - uint32(0x20)) < uint32(0x5f)))
}

func X__isprint_l(tls *TLS, c int32, l locale_t) int32 { /* isprint.c:9:5: */
	return Xisprint(tls, c)
}

func Xisspace(tls *TLS, c int32) int32 { /* isspace.c:4:5: */
	return (Bool32((c == ' ') || ((uint32(c) - uint32('\t')) < uint32(5))))
}

func X__isspace_l(tls *TLS, c int32, l locale_t) int32 { /* isspace.c:9:5: */
	return Xisspace(tls, c)
}

func Xisxdigit(tls *TLS, c int32) int32 { /* isxdigit.c:3:5: */
	return (Bool32((func() int32 {
		if 0 != 0 {
			return Xisdigit(tls, c)
		}
		return (Bool32(((uint32(c)) - uint32('0')) < uint32(10)))
	}() != 0) || (((uint32(c) | uint32(32)) - uint32('a')) < uint32(6))))
}

func X__isxdigit_l(tls *TLS, c int32, l locale_t) int32 { /* isxdigit.c:8:5: */
	return Xisxdigit(tls, c)
}

type div_t = struct {
	quot int32
	rem  int32
} /* stdlib.h:62:35 */
type ldiv_t = struct {
	quot int32
	rem  int32
} /* stdlib.h:63:36 */
type lldiv_t = struct {
	quot int64
	rem  int64
} /* stdlib.h:64:41 */

type ssize_t = int32 /* alltypes.h:88:15 */

type intptr_t = int32 /* alltypes.h:93:15 */

type off_t = int64 /* alltypes.h:185:16 */

type pid_t = int32 /* alltypes.h:258:13 */

type uid_t = uint32 /* alltypes.h:268:18 */

type gid_t = uint32 /* alltypes.h:273:18 */

type useconds_t = uint32 /* alltypes.h:283:18 */

func X__putenv(tls *TLS, s uintptr, l size_t, r uintptr) int32 { /* putenv.c:8:5: */
	var i size_t
	var newenv uintptr
	var tmp uintptr
	//TODO for (char **e = __environ; *e; e++, i++)
	var e uintptr
	i = size_t(0)
	if !(Environ() != 0) {
		goto __1
	}
	//TODO for (char **e = __environ; *e; e++, i++)
	e = Environ()
__2:
	if !(*(*uintptr)(unsafe.Pointer(e)) != 0) {
		goto __4
	}
	if !(!(Xstrncmp(tls, s, *(*uintptr)(unsafe.Pointer(e)), (l+size_t(1))) != 0)) {
		goto __5
	}
	tmp = *(*uintptr)(unsafe.Pointer(e))
	*(*uintptr)(unsafe.Pointer(e)) = s
	X__env_rm_add(tls, tmp, r)
	return 0
__5:
	;
	goto __3
__3:
	e += 4
	i++
	goto __2
	goto __4
__4:
	;
__1:
	;
	if !(Environ() == oldenv) {
		goto __6
	}
	newenv = Xrealloc(tls, oldenv, (uint32(unsafe.Sizeof(uintptr(0))) * (i + size_t(2))))
	if !(!(newenv != 0)) {
		goto __8
	}
	goto oom
__8:
	;
	goto __7
__6:
	newenv = Xmalloc(tls, (uint32(unsafe.Sizeof(uintptr(0))) * (i + size_t(2))))
	if !(!(newenv != 0)) {
		goto __9
	}
	goto oom
__9:
	;
	if !(i != 0) {
		goto __10
	}
	Xmemcpy(tls, newenv, Environ(), (uint32(unsafe.Sizeof(uintptr(0))) * i))
__10:
	;
	Xfree(tls, oldenv)
__7:
	;
	*(*uintptr)(unsafe.Pointer(newenv + uintptr(i)*4)) = s
	*(*uintptr)(unsafe.Pointer(newenv + uintptr((i+size_t(1)))*4)) = uintptr(0)
	*(*uintptr)(unsafe.Pointer(EnvironP())) = AssignPtrUintptr(uintptr(unsafe.Pointer(&oldenv)), newenv)
	if !(r != 0) {
		goto __11
	}
	X__env_rm_add(tls, uintptr(0), r)
__11:
	;
	return 0
oom:
	Xfree(tls, r)
	return -1
}

var oldenv uintptr /* putenv.c:22:14: */

func Xputenv(tls *TLS, s uintptr) int32 { /* putenv.c:43:5: */
	var l size_t = (size_t((int32(X__strchrnul(tls, s, '=')) - int32(s)) / 1))
	if !(l != 0) || !(int32(*(*int8)(unsafe.Pointer(s + uintptr(l)))) != 0) {
		return Xunsetenv(tls, s)
	}
	return X__putenv(tls, s, l, uintptr(0))
}

func X__env_rm_add(tls *TLS, old uintptr, new uintptr) { /* setenv.c:5:6: */
	//TODO for (size_t i=0; i < env_alloced_n; i++)
	var i size_t = size_t(0)
	for ; i < env_alloced_n; i++ {
		if *(*uintptr)(unsafe.Pointer(env_alloced + uintptr(i)*4)) == old {
			*(*uintptr)(unsafe.Pointer(env_alloced + uintptr(i)*4)) = new
			Xfree(tls, old)
			return
		} else if !(int32(*(*uintptr)(unsafe.Pointer(env_alloced + uintptr(i)*4))) != 0) && (new != 0) {
			*(*uintptr)(unsafe.Pointer(env_alloced + uintptr(i)*4)) = new
			new = uintptr(0)
		}
	}
	if !(new != 0) {
		return
	}
	var t uintptr = Xrealloc(tls, env_alloced, (uint32(unsafe.Sizeof(uintptr(0))) * (env_alloced_n + size_t(1))))
	if !(t != 0) {
		return
	}
	*(*uintptr)(unsafe.Pointer((AssignPtrUintptr(uintptr(unsafe.Pointer(&env_alloced)), t)) + uintptr(PostIncUint32(&env_alloced_n, 1))*4)) = new
}

var env_alloced uintptr  /* setenv.c:7:14: */
var env_alloced_n size_t /* setenv.c:8:16: */

func Xsetenv(tls *TLS, var1 uintptr, value uintptr, overwrite int32) int32 { /* setenv.c:26:5: */
	var s uintptr
	var l1 size_t
	var l2 size_t

	if (!(var1 != 0) || !(int32(AssignUint32(&l1, (size_t((int32(X__strchrnul(tls, var1, '='))-int32(var1))/1)))) != 0)) || (*(*int8)(unsafe.Pointer(var1 + uintptr(l1))) != 0) {
		(*(*int32)(unsafe.Pointer(X___errno_location(tls)))) = 22
		return -1
	}
	if !(overwrite != 0) && (Xgetenv(tls, var1) != 0) {
		return 0
	}

	l2 = Xstrlen(tls, value)
	s = Xmalloc(tls, ((l1 + l2) + size_t(2)))
	if !(s != 0) {
		return -1
	}
	Xmemcpy(tls, s, var1, l1)
	*(*int8)(unsafe.Pointer(s + uintptr(l1))) = int8('=')
	Xmemcpy(tls, ((s + uintptr(l1)) + uintptr(1)), value, (l2 + size_t(1)))
	return X__putenv(tls, s, l1, s)
}

func Xunsetenv(tls *TLS, name uintptr) int32 { /* unsetenv.c:9:5: */
	var l size_t = (size_t((int32(X__strchrnul(tls, name, '=')) - int32(name)) / 1))
	if !(l != 0) || (*(*int8)(unsafe.Pointer(name + uintptr(l))) != 0) {
		(*(*int32)(unsafe.Pointer(X___errno_location(tls)))) = 22
		return -1
	}
	if Environ() != 0 {
		var e uintptr = Environ()
		var eo uintptr = e
		for ; *(*uintptr)(unsafe.Pointer(e)) != 0; e += 4 {
			//TODO if (!strncmp(name, *e, l) && l[*e] == '=')
			if !(Xstrncmp(tls, name, *(*uintptr)(unsafe.Pointer(e)), l) != 0) && (int32(*(*int8)(unsafe.Pointer((*(*uintptr)(unsafe.Pointer(e))) + uintptr(l)))) == '=') {
				X__env_rm_add(tls, *(*uintptr)(unsafe.Pointer(e)), uintptr(0))
			} else if eo != e {
				*(*uintptr)(unsafe.Pointer(PostIncUintptr(&eo, 4))) = *(*uintptr)(unsafe.Pointer(e))
			} else {
				eo += 4
			}
		}
		if eo != e {
			*(*uintptr)(unsafe.Pointer(eo)) = uintptr(0)
		}
	}
	return 0
}

type uintptr_t = uint32 /* alltypes.h:78:24 */

type int8_t = int8 /* alltypes.h:119:25 */

type int16_t = int16 /* alltypes.h:124:25 */

type int32_t = int32 /* alltypes.h:129:25 */

type int64_t = int64 /* alltypes.h:134:25 */

type intmax_t = int64 /* alltypes.h:139:25 */

type uint8_t = uint8 /* alltypes.h:144:25 */

type uint16_t = uint16 /* alltypes.h:149:25 */

type uint32_t = uint32 /* alltypes.h:154:25 */

type uint64_t = uint64 /* alltypes.h:159:25 */

type uintmax_t = uint64 /* alltypes.h:169:25 */

type int_fast8_t = int8_t   /* stdint.h:22:16 */
type int_fast64_t = int64_t /* stdint.h:23:17 */

type int_least8_t = int8_t   /* stdint.h:25:17 */
type int_least16_t = int16_t /* stdint.h:26:17 */
type int_least32_t = int32_t /* stdint.h:27:17 */
type int_least64_t = int64_t /* stdint.h:28:17 */

type uint_fast8_t = uint8_t   /* stdint.h:30:17 */
type uint_fast64_t = uint64_t /* stdint.h:31:18 */

type uint_least8_t = uint8_t   /* stdint.h:33:18 */
type uint_least16_t = uint16_t /* stdint.h:34:18 */
type uint_least32_t = uint32_t /* stdint.h:35:18 */
type uint_least64_t = uint64_t /* stdint.h:36:18 */

type int_fast16_t = int32_t   /* stdint.h:1:17 */
type int_fast32_t = int32_t   /* stdint.h:2:17 */
type uint_fast16_t = uint32_t /* stdint.h:3:18 */
type uint_fast32_t = uint32_t /* stdint.h:4:18 */

// Support signed or unsigned plain-char

// Implementation choices...

// Arbitrary numbers...

// POSIX/SUS requirements follow. These numbers come directly
// from SUS and have nothing to do with the host system.

func X__strchrnul(tls *TLS, s uintptr, c int32) uintptr { /* strchrnul.c:10:6: */
	c = int32(uint8(c))
	if !(c != 0) {
		return (s + uintptr(Xstrlen(tls, s)))
	}
	var w uintptr
	for ; (uintptr_t(s) % (uintptr_t(unsafe.Sizeof(size_t(0))))) != 0; s++ {
		if !(int32(*(*int8)(unsafe.Pointer(s))) != 0) || (int32(*(*uint8)(unsafe.Pointer(s))) == c) {
			return s
		}
	}
	var k size_t = ((Uint32(Uint32FromInt32(-1)) / size_t(255)) * size_t(c))
	for w = s; !(((((*(*uint32)(unsafe.Pointer(w))) - (Uint32(Uint32FromInt32(-1)) / size_t(255))) & ^(*(*uint32)(unsafe.Pointer(w)))) & ((Uint32(Uint32FromInt32(-1)) / size_t(255)) * (size_t((255 / 2) + 1)))) != 0) && !(((((*(*uint32)(unsafe.Pointer(w)) ^ k) - (Uint32(Uint32FromInt32(-1)) / size_t(255))) & ^(*(*uint32)(unsafe.Pointer(w)) ^ k)) & ((Uint32(Uint32FromInt32(-1)) / size_t(255)) * (size_t((255 / 2) + 1)))) != 0); w += 4 {
	}
	s = w
	for ; (*(*int8)(unsafe.Pointer(s)) != 0) && (int32(*(*uint8)(unsafe.Pointer(s))) != c); s++ {
	}
	return s
}
