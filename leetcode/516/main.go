package main

import (
	"fmt"
)

var m map[string]int

func longestPalindromeSubseq(s string) int {
	if m == nil {
		m = make(map[string]int)
	}
	if len(s) <= 1 {
		return len(s)
	}

	if cnt, ok := m[s]; ok {
		return cnt
	}

	cnt := 0
	if s[0] == s[len(s)-1] {
		cnt = 2 + longestPalindromeSubseq(s[1:len(s)-1])

	} else {
		l := longestPalindromeSubseq(s[1:])
		r := longestPalindromeSubseq(s[:len(s)-1])
		if l > r {
			cnt = l
		} else {
			cnt = r
		}

	}
	m[s] = cnt

	return cnt

}

func main() {
	rst := longestPalindromeSubseq("bbba;lsdfjh;oiwehgf;alsdhf;alsdkjfla;sdjfowehfijfal;skjdf;owiaejf;oawejf;oawiejf;oaswiejf;osidjgf;olsdfjg;lskdfjgwseo;irgjfw;soerigyh;osdfhvn.kxcnvb.,xzcmnbvs;dofvhboi;seruvns;eorvnms;ldfghjksl;dfkjg;soedrgj;oisehrg;sldjkf;sljtoiewsr;tgh;ewsorgj;sldxfjg;sldjgf;soleirjfgio;sejrg;osldirjg;sdofljg;lsdfgj;sdlfvjkns;eoirhv;seoriugj'sdl;kfjg;sldjfg;sopeirjfg;sdlfghjv;lsdfvjn,x.cmjnvsel;dfkghs;eorigj;s'eopjgf;sldfkgjvx.d,cfvmnbs;ldfkhgj;waoierujt;woeriujtg;sldrgjh;lsdfjgh;soeritujghhq;woeriujtp0qw3457utpqw3o4;thj3wperoguwq3p0498kw8p[03f4h5p39w74h5fpw934uytngh;w3oeh8ruf;wesoaf83kj4589w3j4fpab")

	fmt.Println(rst)
}
