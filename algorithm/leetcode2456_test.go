package algorithm

import (
	"fmt"
	"testing"
)

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {

	if len(creators) == 0 {
		return [][]string{}
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	maxView := 0
	sumViewMap := make(map[string]int)
	for i := 0; i < len(creators); i++ {
		val, ok := sumViewMap[creators[i]]
		if ok {
			sumViewMap[creators[i]] = val + views[i]
		} else {
			sumViewMap[creators[i]] = views[i]
		}
		maxView = getMax(maxView, sumViewMap[creators[i]])
	}

	maxIdxMap := make(map[string]int)
	for i := 0; i < len(creators); i++ {
		sum := sumViewMap[creators[i]]
		if sum != maxView {
			continue
		}
		idx, ok := maxIdxMap[creators[i]]
		if !ok {
			maxIdxMap[creators[i]] = i
			continue
		}
		// 这里得判断view，所以存索引
		if views[i] > views[idx] {
			maxIdxMap[creators[i]] = i
			continue
		}
		if views[i] == views[idx] && ids[i] < ids[idx] {
			maxIdxMap[creators[i]] = i
		}
	}

	keys := make([]string, 0, len(maxIdxMap))
	for key := range maxIdxMap {
		keys = append(keys, key)
	}
	res := make([][]string, 0, len(maxIdxMap))
	for _, key := range keys {
		idx := maxIdxMap[key]
		res = append(res, []string{key, ids[idx]})
	}

	return res
}

func Test2456(t *testing.T) {
	creators := []string{"u", "ajf", "n", "kkmq", "mwkim", "p", "ktjvr", "ihmh", "oulo", "b", "q", "ofdim", "rqbft", "mdf", "txt", "xyjv", "rlx", "re", "fyd", "dq", "frc", "fag", "xshlj", "z", "gii", "z", "le", "fcvgf", "yqbnk", "vhke", "udvp", "rb", "ppy", "jywvl", "xj", "hb", "ppqsq", "waf", "wpuw", "qg", "rnux", "d", "kxbcl", "yoaqf", "hnphp", "w", "nivm", "hvymz", "xze", "bq", "u", "wbye", "lmqoo", "pc", "q", "t", "jgiy", "guv", "fyc", "ng", "pvlg", "aj", "fhdo", "maeu", "zwfun", "ravm", "yypgx", "cd", "fkzmb", "tvq", "dm", "aphdl", "rbcp", "dtcr", "ehcv", "k", "c", "hc", "tg", "wgin", "mrrr", "glr", "fvxy", "cap", "xjjtq", "hqp", "bn", "t", "bc", "cbbwf", "ztxnz", "xzmw", "wsx", "osim", "m", "cr", "sp", "s", "v", "he", "mhcp", "flz", "owcx", "zzi", "p", "wvvm", "su", "jp", "qf", "icqz", "yy", "k", "jfv", "qscyf", "v", "wj", "nhlsi", "r", "vmd", "nnbca", "u", "s", "r", "uoavb", "qm", "m", "wio", "ernca", "h", "bzrv"}
	ids := []string{"z", "w", "scghn", "mnmvy", "xgnhf", "khuxq", "hei", "wsowq", "yae", "cs", "h", "hyrrv", "vli", "pma", "bxsh", "xmm", "qkimd", "ut", "fj", "xyzw", "scjsj", "y", "k", "c", "qgx", "fgk", "mg", "rmgse", "txsgi", "fzn", "z", "t", "ew", "yi", "wzitv", "tqg", "b", "o", "sesb", "jpw", "u", "rwc", "ermmg", "rjsjw", "qh", "mqf", "ax", "anh", "hanz", "ooors", "mv", "shaca", "doon", "d", "x", "f", "egmiy", "lbfvj", "edrsz", "epwai", "spvwi", "xlh", "eux", "c", "flw", "udo", "bmft", "ohnl", "o", "novqs", "l", "vosc", "nasy", "p", "vk", "cx", "krdo", "zdusc", "pm", "pcc", "ye", "sx", "cjjx", "je", "i", "iywdt", "sd", "kmx", "dfq", "kcq", "zbgjc", "awvkp", "utdq", "wos", "y", "sch", "jmsxr", "aewo", "ngy", "b", "tt", "dfzb", "db", "nzm", "fl", "om", "s", "gmpa", "ie", "yj", "nbey", "v", "yz", "oqf", "glo", "daeig", "wim", "ay", "d", "qsgp", "l", "y", "er", "e", "pz", "wn", "ys", "upvkl", "lzjn", "fjs"}
	views := []int{29, 383, 953, 680, 836, 892, 572, 308, 987, 154, 409, 689, 693, 144, 187, 104, 95, 683, 987, 723, 196, 220, 429, 194, 840, 201, 408, 283, 329, 530, 657, 73, 897, 888, 261, 177, 32, 87, 948, 752, 367, 190, 546, 575, 223, 936, 549, 367, 148, 350, 217, 393, 989, 834, 730, 425, 799, 835, 325, 960, 749, 809, 842, 270, 172, 731, 567, 739, 707, 581, 261, 106, 409, 63, 576, 27, 288, 693, 567, 950, 967, 400, 552, 869, 456, 579, 355, 392, 977, 394, 945, 423, 529, 804, 243, 308, 508, 218, 210, 155, 857, 875, 838, 283, 632, 641, 17, 80, 688, 780, 403, 300, 846, 65, 315, 132, 437, 898, 460, 374, 526, 870, 421, 696, 669, 133, 17, 937, 200, 377}
	res := mostPopularCreator(creators, ids, views)
	fmt.Println(res)
	//for i := 0; i < len(creators); i++ {
	//	if creators[i] == "p" {
	//		fmt.Println(i)
	//		fmt.Println(creators[i], ids[i], views[i])
	//	}
	//}
}
