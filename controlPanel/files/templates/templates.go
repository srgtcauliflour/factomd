package templates

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

type staticFilesFile struct {
	data  string
	mime  string
	mtime time.Time
	// size is the size before compression. If 0, it means the data is uncompressed
	size int
	// hash is a sha256 hash of the file contents. Used for the Etag, and useful for caching
	hash string
}

var staticFiles = map[string]*staticFilesFile{
	"general/footer.html": {
		data:  "{{define \"footer\"}}\n    </body>\n</html>\n{{end}}\n",
		hash:  "239cf4d42e3fd75d2fa4395fb31f7c6679f9c248afe058dc90d9b37280bc5aa2",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755049, 0),
		size:  0,
	},
	"general/header.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\x9cUM\x8f\xe36\f\xbdϯ`u\xe9%\xb2[\x14-\x8a\x998\xa7b\x8b^zY\xa0ׂ\xb1h[]Y\xf2J\x943\xd9\xc1\xfe\xf7\xd2vҸ\x13\xcf.\xda\x00\x89\xad\x8f\xf7\xf4H>1//\x86\x1a\xeb\tTGh(\xaaϟ\x1f\xf6ߘP\xf3y \xe8\xb8w\x87\x87\xfd\xf4\x80\xdaaJ\x95\xf2A\xff\x95\x148\xf4m\xa5\xc8+06V\xcaqT\x87\a\x90\xcf~\xe2Y^\xe7aO\x8cPw\x18\x13q\xa527\xfag\xf5z\xb9c\x1e4}\xccv\xacԳΨ\xeb\xd0\x0f\xc8\xf6\xe8HA\x1d<\x93\x17\xac\xa5\x8aLKwh\x8f=Uj\xb4t\x1aB\xe4\x15\xe0d\rw\x95\xa1\xd1֤\xe7\xc1\x0e\xac\xb7l\xd1\xe9T\xa3\xa3\xea\xfb\xe2\xbb5\x9d\xb3\xfe\x03Dr\x95J\x9dPՙ\xc1\n\x9b\x82)\x19\"\xa0ǖ\xcag\xbd\xccu\x91\x9ai\xae-\xa7q\xd9\xe08=\v\xf9YS\xb2eG\x87߃!\xf8æ\x8c\xce~\xa2\b\x1a\xdea͡\x87w!{#\x91\x06\xbf/\x97\xad\x9bj\xf8\xec(uD|=\xb6N\xa9l\xfe\xc1\x16\xbd\xf5\x85L\xa9\xff\x84\xf6\xac\xf1D)\xf4\xf4\u007f\xf08\f+Ⱦ\xbc\x95}\u007f\f\xe6|c\xba\xb9\x82\xe2\x8a\xdf\xd8\xf1j\xa9\x18N\xab\x93_\xaf:\x8c-\xe9\x1f\xa0'cs\xaf\u007f\x94\xfa\xba\xdc\xfb\x04.\xb4\xe1\x15n\xc6\xe2Ee\xa9\x0e{)\x0f\xa4X/uj\xe6\x9c\xeb[\xde\xfe\x94*\xdbO\x92\btE\x1a[\x05\xe8\xc46w\xa5\x11\x9e\x12_\t,E\xe1\xd74\xfft\xaf9\x11ƺ\x83\x8b\x92e\xa4'â\xdc\xc1\xb8\x15͊\xd5\xfa!\xb3nc\xc8\xc3\xc6\xceyw\x1a\xd0ol\xd7\x0e\x8f\xe4\xa6|\\\x17\x1b\x14\x11\x17\x01\x12w\xb4\xa8;k\f\xf9Jq\xcc4\x85l\xe5;\xf1\xbdq\xd4L\u007f\xb9\x19W\x9e\xe5*~T[\x12\x1aK\xce(\xb0f:|\x15\xbe\x82\xc1aM]pb\x90J\xbd_2$E$\xdbv\xbc\x034&RJ;\xe0\x88>\tR*\x02\xbf\xfd\xb2\x03\xe2\xba(\xdeJ\xc4v\xda\xf413O\x05\xdd\x04݇\x95\x8f\xbd\xe5\r\xcd\xfa\xbar9b\xa1\x05\x83\xf1\x83\x82\x11]\x16\xf0\xaf\xe1-m\xf7\u07b9Mö\x03\xe6;X\xa9!$;\xc5\xff\x88\xc7$\x8ebz\x9a\xdbڣ\xf5\x1dE\xd1\xf3\xf5d\xfc;\f\x8a1D\x90V\xe8\x82\x04-\x1d1J\xe8\xf4̺\x96\x06*~\xbc\x9e\x1bF\x8a\x8d\v\xa7\xc7\xc5#O\xc6&)\xda\xf9\xd1\aOO\xea\xf0>\xc4x\xde\xc19\xe4xu\xb8\xb1\xc6\u007f\xcb\xd2@8G\x0f\xe8\xcf\xf2\x9a\xb2\xe3T|1\xfc/]\xb2\xd5pi7S?yy!o\xe4?\xeb\xef\x00\x00\x00\xff\xffJ\x11\x01\u007f\xc8\x06\x00\x00",
		hash:  "5669ff839696736c7908c0f37f39d80497b8f2669e87a5cb5226c256ace710d2",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1476234260, 0),
		size:  1736,
	},
	"general/scripts.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\x8c\xcdK\n\xc3 \x10\xc6\xf1}O!\xee\x13/\x90\xe6.\x83\x0f:BF\xeb\x8c} \u07bd\x01w]\x04g7\xf0\xfb\xf3\xb5\xe6|@\xf2J\xb3-\x98\x85u\xef7u\xde6~\xc5\xc5\xdeud\xf3\xf2\xe4R1\xf1Y}\xf9\xae\x91\xf5\xbe\x99A\xf6+\xff~\x80,H\xb9\xca|\x13R%\a\x82\x89\xd6\x03i\xa2\x83\x9c'T\x00+\xe9p\vD\xf8\xfc\xf1\xd6\xce\xe5\xde\u007f\x01\x00\x00\xff\xff\xb1\x1b \xa8\r\x01\x00\x00",
		hash:  "e084eeb01388db75a5076e1e487fc761e3b7ece8ee41f6a104896b630431b97c",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755049, 0),
		size:  269,
	},
	"index/controlPanelScripts.html": {
		data:  "{{define \"controlPanelScripts\"}}\n\t<script src=\"js/controlPanel.js\"></script>\n{{end}}",
		hash:  "5fc312858cfbfba3160a20303376141a6fcfb759d1291c25dcb27fe9aecffae0",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755049, 0),
		size:  0,
	},
	"index/datadump.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xecW]o\x9b0\x14}N~\x85\xe7={\xa8_y\"H\xfbx\x99\xb4lӲ?\xe0\xc2M\xb1j\f\xb2M\xd6(\xca\u007fߵ\x81@h\xba\xa5$բmy(\x96/\xf7\\\xfb\x9cs\xb1\xbb^'\xb0\x10\n\bM\xb8\xe5I\x99\x15t\xb3\x19\x87\x06b+rED2\xf5\x81\x0f.@bɍ\x99\xd2T$@\xa31\xc1_\x98\x88e3\xad\xf3\x1f8;\xeaOǹ,3e\x9a\x90\x0f\xa7\x17\xd1wЙP\\\x12\aODV\xe4چ\x01\x06ƣ\xd1(,e\x93n\xf9\xad\xa1\xfe%\xe6\x86~E\xf0\xc0\xb3B\x82\x9f\xa0>a\x14J\xd1\xcd`VX\tD\x18\xc6q#K\\n\xc8I\xaaa1\xa5\xaf\xdd&/(\xe1Zpf@\xe2N\x011\xad.\xf1\xady\x99e\\\xaf\u0080Ga \xc5/\xb0\xfb\x88\x974\xfa\xaa\xf3\x18\x8c!\x9f\x84\xb1\x03\x10\xae\x1c\x82PvƋ\x01\xd9\u05f8x\xd0K\xd0f@\xf2\r\x8d\xde\xe7JU\xa2\x0f\x01\x98\xd0\xe8]\x19߃%\x93\x9d\xec0(e5\xe88\xc2\xc3Ĺ\xb2\xa0lG\xdafj\xbf\xbe\xfd\xfc\x82+\x90\x1d\x81+\xabzi\xab\x8c\xbe\x89\x88\v\xb2\x03\xed\x84ٯ\x18\xebl\xbdI&\x87\xb8k\x9e\xa2\x99\x9frX\xea\x8d^s\xc4ض\xde\xc0Z\xdf8v\x1d\xfe\xd9\x15\xad%\xbe\xa6\xaeaǓ̱\xe1u\xbbS\x91\xdd\xf9\xf8\xa2\x94\xd2\xc4\x1a@\xb1\xbcpN\xd8v<\xae\x03\xbb\xd8\x02\xdb\xf3\x8a\xd1\xf1\x94\"D\xd0\xc6ޘ\xe5\x1d\xae3\xc0YW\x84t~\x15\xaf\x16\x1e,\xd7\xc0I\"\f\xbf\x95\x90\x10S\x80\x94q\n\xf1=\xae\x83K\x03\xf40\xb5+\xaa\xb1V\x03٥\xf4te<˝\"\r\xc9Hm\xed\xcf\xcep\xafU[\xac\xcb\ue9f0\xe2\xe48\xa7\x1e\xe1\x9c\xfd\x1e=o;\x9d\xb1\xb4W\xff\xa5\xfd[\xa5\xbd>ѱ2Pѷ\xa5M\x9f\x90ԅr-\xac\x80\xde\xc9\xfd\xbbj\xfd\x1a\x1f\x13w G\xfeq\x02\xb8\xd9\xeas\xee.\x8a\xb3\x15q\x83\xf3\xb5\xddNÞ\xd2\u007f^\xb4\xc7\x06\x1cZ\xa2\x05\xae\x95z\x01\xe4F\xb4\x13u\xcd͟\xed\x1a\xbc\xd4>\xe7S\xf8L\x83\xcf\xfd\xd5\xe3\x8bMA\xff\x83殹}\t\x13\xce\xfbw\xbac,8y|&oq\xfb%Z\xd8\xed\xa8\x19\xd4\xcf0\xa8\xff5\x8e\xc6\xeb5\xa8d\xb3\xf9\x19\x00\x00\xff\xff[-҇B\x0f\x00\x00",
		hash:  "ead616476b3fc70b5e4b97a04398a5920e49e52195d4180c79c06b3383993c22",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755049, 0),
		size:  3906,
	},
	"index/index.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xfft\x8f\xc1\n\xc20\f\x86\xcf\xfa\x14u\xf7\xfa\x04\xe2A\xf0>\xd0\x17\bM\xa6\x856)m&\x8e\xd2wW\xe6At\xec\x9a\xff\xfb\x92\xfc\xb5\"\r\x9e\xc9t\x9e\x91\x9e=ܨkm\xbb\xa9U)\xa6\x00\xfaN\xee\x04Hy\x1e\x1fv֚\x93\xe0d\xac=\xfeR\xb3\xcf\xf0X\xe8A\x1c\x84\xab\xa4\xce\xec\xff#\xcd\xc0\x05\x9cz\xe12\xc6\byZ\xd8\b\n8\xc6\xf4=\u007ff\\y\xa1\xb8쓖\xc5\x0e'\xacYB\x0fL\xe1\xb2\xc2\f\"\xfa)Y+1\xb6\xf6\n\x00\x00\xff\xff)\xb2x\xeb\x1a\x01\x00\x00",
		hash:  "cc3a3da9da68933167db530679efab736e844d7d525da9b7f55fe484cfc74a3f",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755049, 0),
		size:  282,
	},
	"index/indexnav.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\x84\x90\xb1n\xf30\f\x84w?\x05\xc1]\xf0\xf4/\u007f\x15M]:\xa4(\xd0'`,& @K\x81D\xbb.\f\xbf{\x15\xa3\x01\x9a.\xd5 \x10:\xdd}\a\xae+D>Kb@I\x91\x97D3n[\xe7\xa3\xcc0(\xd5z\xc0\x92?\x10\xaa}*\x1fp\xa4r\x91\xe4N\xd9,\x8f\xff\xe1\xdfuy\xc2\xd0A;~һ\xc1\xe8T\xe1v\xb9!'+Yݕ\x12+B$#\xb7\xab\x12\x0f\xc8\v\x8dW\xe5\xfd\xe1;d\x0fR\xf9\x19\xe4LL\x19\xa4:\x1aLf\xc6\xdd{\xef\xeaF\x92\x84\xc1\x13P\x11r\x95\x95\a\xe3\xf6\xc1\xca\xc4\x18\x8eM\x85w#\x9b*\xbcх}O\xc1\xf7*\u007f\xd1\xda\xf8\xbb\xfb#5\x17\xbeQñ\r\xf0\xccF\xa2\x1c\xe15G\x86\x97t\xcee$\x93\x9c\x1eq\xbe\x9f4t\xbeo\xab\rݺr\x8a\xdb\xf6\x15\x00\x00\xff\xff\xa7\xffC\x19\u007f\x01\x00\x00",
		hash:  "dbe5cd4fb16f7f214469c016e8613fc72fd7f0bba689d5a13a5a99fac9615dc8",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1476137647, 0),
		size:  383,
	},
	"index/localTop.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xdcWQo\"7\x10~\xbf_Ẫ\x94H\xf5\x01i\xd4F)\xacDRݵ\xd2\xf5zj\xd2J}4\xeb\x01\xacx\xed\xad\xed%A\x88\xffޱw\x17\x12\x0e\u0605\x1c\xf7\xd0<$`\xcf|3\xf3\xcd米X\b\x18K\r\x84*\x93ruor\xba\\\xbe!\xd5O\xdfA\xea\xa5\xd1D\x8aAi@\x93\xd5f4\x10rFRŝ\x1bPk\x1e7v7-R\xa3\x8aL\xbb-Ve\xb0\x8c+\x95\xf4]\xceˀ\x0e\xec\f,s\x9e\xfb\x02\x9d\xfa\x9d\xb0\x13\xfeD\xbb\xed\x18\xd3\x1eq~\xae`@\x1f\xa5\xf0\xd3\xeb^\xb7\xfb\xdd\xcf4\xf9\xc7\x14\x96|4\x02\xea(\xb3\x8b\xb7=\xd2\xe7dja<\xa0\xdf\xd2\x18p\xccSo2\xe6\x80\xdbtʔ\xd4\x0f49\xfb+\x9fX\x8e\x8e\xc3\x19\x97\x8a\x8f\x14\x9c\xf7;|\x95E\tW\xc7\x1c+\xc3\xfd\xb5\x95\x93\xa9Ǡ\xef\xa5'7\x85T\xe2\x9a,\x16o\x97˕Og\xdaK\xc8\xf6\xfc\xbfa\xacJ\x91\x18\x9d*\x99>\f\xa8\x86'\x1fr?;G\x12x\xf2\x11\xbf\xc6Z\x9e\xe7\xf1}]\xd8ma-h\u007f\xbd\xa61-W\x98F\x17\xa6\x8bl\x04\x96&\xdd\r6\tc[z\xd7\xc1\xe6m4|\xcb\xd2A\x1aP\xdcN\x80\xfd@2\x10\xb2\xc8\xd8%\x89\xf1Y\xef\x824\xa8\xe3\x19F\x06\xde\xcat\x87a4\xc6F\x81\"cc\x91=,\xfbW\b-\xa9dp\x83B~ \xe5\xd2u\xbf\x13M\xf7@I\x9d\x17\x9e\xf8y\x8e\xfd\xf5H})\x95g\xa8D\xf3\f^\xae\b\xe9\x82R\xd0\xce\xdb\x02(\x99qU\xa0\t\xdbU\xdb社\xbal7\xd7\xe9;i\x1dV\x1du\u007f\x87\xdf\xc9]<J\xe4\xac\xe7<\xc9\x11\xf3\xbcE\xfd!\x81x\x1aW\x80u>\xb95\x13\v\xceQbMP\u007f\xfd}\xc4-%\x9e\x8f\xa4\x16\xf04\xa0]J\xb8\x95\x9cE\x12\xb4y\x1cЋ\x17K\x99\xd4\x1bF\x81\xe6\x01ţKr\xb0)j\xf7\x859\u007f\x8a{{x(\xa7I\xd0\xffF\xa6\f)D\xf5\xbf\x1c\x11$̈\x06\xb4\x88\x98G\x1e<8/\xf5\x84n\xc7fQ\"I\x80\x8c\x94\x83 g]bƤ\x8bc#oH\xb9<\x92\xbb[\xb1G&'R\xd0\x1d\xa4F\x8bm\x12\xba\xd0\xe2(\tU\x88_ICW\x97_GBW\x97-\x15\xf4\xbf\x14M\xe3\x05\xb0˺\x9c\xfd?6\x8c\xfe\x9d\x02\x8d\xef\x831\x88\xd4\x14\xd8\xe0\xe4\x1d\b\xb0܃hTd\xc3p\xdf\x00\xae\x06\xfc\xe6\xea\x81C\xbe\x05맢\x88\x175E\xc3BH\xffe\xe8Y\x81V\xf4\x9c\x8a\x90\xc3\x05\xbck\xf9\xb3W\xc8U\xfd\n\xf9\xe9䯐\x1c\xc0~\x90\xe16^?̼\xf1\\}\u008d۲9\xd5Q&\xb7F\xeb\xf2\xdd\xedZ\fW\x1f8\x8fx\xeb\x18\xfb\xf9\xf6S\xe0\xa2E\xf7\xbdm6\xaa\x00W\xf1\x99\xcci\xf2\xdb'\x94N6!q8\x86I\xbb\x1a\xf7\xce\xd8py\xb2\xb0;\x95\x02\xe8sG\x16v\xc3\x16%\x1d$\xc3O[\x87O\xca[\xe90\x9f\x83\xac\xd7y\x8a\x02g\f\xf6\x86&\xbfT\x9f\x8e(\xb6\x069\xb6d\xfc7!\x94Pg\x10=\xb7\xbd\xe0\x1b\xabq\xe1jL\xee\xf0\xf7\x11U\x04\xe7㛶Ʊ\x90\x82\x9c\x01\xbe3\xfe\xac>\x1d\x91L\r\xf2\n\x15\r\xcbC\xd7\xce\t\xadl\xd3`kq\xd2\xfa~dļ\x11\xa8\x85\x91\x1f\x1b\xe3\xbf\xe8\xb1\x16\xc9\xfd\x1f\xf7\xc3\x0f\x9d\xe1\xdf\xef1\x05\xd1\xde\xed \xebU\v\xff-\xb8\x92~N\x93S\a+pFu\xc3\x13\xeb\xf7\x9b\xf3ý\x85y\xd4G\xfa\xb7̵\x95\xb4\xf6w\x1b\r½\xf0\xda{sc\to\xa8\xf2bJ\xde,\x16\xa0\xc5r\xf9_\x00\x00\x00\xff\xff3˗\xe5@\x12\x00\x00",
		hash:  "dba5902c4c554ab8697889835fa1185848f5bc59222a986e3d433abef0d0517a",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1476131421, 0),
		size:  4672,
	},
	"index/transactionsummary.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xc4Vak\xdb<\x10\xfe\x9c\xfc\n\xa1\xc2˻\x0f\xc6\xdb>n\x8aaMZ\x1a\xd61\x18\xfd\x03\x8a\xa5̢\xb2\x1c\xa4sW\x13\xf2\xdfw\x92c\xe3&q\xe2t45\x84\xc8w\xf7H\xcf\xe9ѝL\xd6k!\x97\xcaHB\xc1r\xe3x\n\xaa0\xae\xccsn+\xbaٌ\x99\x93\xc1D\x94\x98\xbc\b\xa1ɘ\xe0Äz\"\xa9\xe6\xceM\xa8-\xfel\xad\xbb\x9e\xb4\xd0e\xdebڈ\xecS\xf2Й\x92\xfc\xc7\xf3\xd5Wrc\xc0*\xe9X\x8c\xee\xf1h4b\xa5n\xe6\x01\xbep\x94\b\x0e<\xf2\xc3@J>#J\xcb`\xa0\x010bZu\x11\x11(В(\x17\xf9\x85\x9e$M\x18'\x99\x95\xcb\t\xbdZqs\x8b\xd6B\t\x9c\x98[\xc5#'5\xa6,C\xba%\xc66n\xe606\xache*\rD\xcb\xda\x11A\x01\\\xd3\xe4\xff\x8f\x1fX\xecc\x12\x16s\xfciu\x84\xcc\x0e\x85m\xca4\x99f\\\x99ؿVdZ\xe4\xb9\x02G\xf6\x16\x96\xde}lY\xb2\xf3\f\xa10S8;\x14\xb6\xba\xd6E\xfaH\x93{\ue034F\x12\xac$\xda'#\x9a\x90hQ\x03\x0fm\x02\x8bK]\x0f:\x87\"PI\v\x038MG\xd4\xc6tX\xd9]<.%uG\xda@\xad+\xea\x81\xdd@\xa0?\x0e\xf5\x81~\xbeW\x0e\x0eDՑ\x99\xe4\u2c2f\xf6\xdb~\xe7v\x82\xee\t'\xf3\x19\x8b\xd1t\x1a\xe3\xb5%s\xb3*a\x18\xe0\xaa\x0e&߄\xb0\xd29_=\xc3`?K8\x03\x87ޞ\x8c=\xaew\xaf\x18,\nQ5\xbe>|7\xe6\x85\xc3˵\x95?F\xfd\x8f\x9d\x84V\xff\xb6\xa2\xde[\xfe\xba\x96\xef\xb8ˆI\x12:\xc0\xe0\x83r3\xc56\xe1\xe0MU\xfbw\xbd\xf6b\x8ek\xb7ۊ\xdeYB\x82\xb7\x97oi\x13\xfa\xb9\xa7-\xceͲ\xb09\xf7%~\x81\xfayU\x16\"\xf9.\xab\x1f\xbf\xbe\xe0:\xe2dl\xd8\xd9\xd9u@\x84k¿\x87\xeb.\xc7\xeb\x91\xdb4\x8b\xb42\x8f\x94@\xb5\x92\x13*\xda\xce\xef[\xfe\xb1\xf9\xfb\xf3\x1f\x9c\xc65\xee\x049?\x17\x0fk\xf2yk\x8a\xb7\xa5֡\xe2\xcfb\xe8Q\x1et\x01\x82\x0f*\x97\x0e\xf0r=o\v\xbd\xca-\xf4\x024\xeb⺓\xeaw\x06\xe73\xadq\xaf\xa7y\xba\xc3\xed;\x9a۩\x1d5\x83\xed?~\x1b՟\xd3\xc9x\xbd\x96Fl6\u007f\x03\x00\x00\xff\xfff\x97\xc7~\x81\v\x00\x00",
		hash:  "8a5c3afc8ccb05195c89284ff19c6370e4b5e7ab7889450aa99b31f429396e16",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755049, 0),
		size:  2945,
	},
	"searchresults/tools.html": {
		data:  "{{define \"tools\"}}\n\t<script src=\"js/searches/tools.js\"></script>\n\t<link rel=\"stylesheet\" href=\"css/searches.css\">\n{{end}}",
		hash:  "a98871ca472c8a4b10c9916ad166c30071ff142d7e9e52290254dd7f836e036a",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755049, 0),
		size:  0,
	},
	"searchresults/type/EC.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xa4\x93An\x830\x10E\xd7\xe4\x14.{\x8a\xba\xad\x9cHI\xca\rz\x01\aO\x14K\xc6F\xf6$-\xb2|\xf7\x0e\xe0\xaaT\x85\x14\xb5\xac\xf0\xff\xf3\xe1\xc9\xfa\x13\x82\x84\xb32\xc0\xf2\xea\x98Ǹ\xc9B@hZ-\x90\xa4\v\b\tn\x90\xf9CQ\xb0\x83\x95\x1d+\x8a\x1d\x1d=Ԩ\xacaJnsxo\xb5u4HFƥ\xba\xb1Z\vﷹ\xb3o\x83\xf6M\xac\xad\xbe6ƏF\xc6/O\xbb\xbd\x94\x0e\xbcg/\x02\x05/Iذ\x99\x87\xa38i\x98\xf72\x8e'b\x9b7\xa7\x9fp\xbf\x8d\xa49\xf9\t\xf5\xccK:\xac\r\x85\xf0\x98r1\xae\t\xd2\xcc\x02\x11\xdd̲\x95\r\x80\x8d\xbd\x1ad\xfb\x9bP\xba\xbf\x99;\xa4)At\a\xa1\x85\xa9!FV\x19t\x1d;:\x90\n\xfd\xbd\xe8\u007f\x18_\xbbv\x05\xd7\x14\xe5\x8f$d-\x17\x80̱:\xfd\xffJ\xaa\xe2\xd0\xd3\xf4\xc2\xcbT\xe5]*ye\xe4Wч\xfct%|\xedT\x8b\xbe߉\x1f\x1eZ\xab睳\xb58.R\b`d\x8c\x1f\x01\x00\x00\xff\xff\x9d\x93a\x14w\x03\x00\x00",
		hash:  "5decf7ae2b1cc77506b4c190c36b0a297816c2dbbf15f1cfd7c93b2ad5c606f9",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755049, 0),
		size:  887,
	},
	"searchresults/type/FA.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xa4SQj\xeb0\x10\xfcvN\xa1\xe7\u007f?\xd3ߢ\x18\x1cڜ\xa0\x17P\xa4\r\x11Ȓ\x916i\x83\xd0ݻ\xb6Up\xa9\x9d\xa6\xad\xbf\xac\x9d\x9de\x18fbTp\xd4\x16X\xb9o˔6E\x8c\b]o\x04\xd2\xe8\x04B\x81\x1f\xc7\xfc_U\xb1\x9dSWVU\r=\x03H\xd4\xce2\xad\xb6%\xbc\xf5\xc6yZ$\xa0\xe0J_\x984\"\x84m\xe9\xdd\xeb8\xfb4\x94Μ;\x1b&\xa0ে\xa6U\xcaC\b\xecI\xa0\xe05\r6l\xe1\xe3(\x0e\x06\x96\xb1\x82ぴ-\x83\xf3\x13\xfe\xbb\x95\xbc\xa7>D=\xf2\x9a\x1e\xf7\x92b\xfc\x9fy)\xddC\xa4\x9d\x15E\xe4\xcc:T\x8c\x02;w\xb6\xc8ڋ\xd0fp\xe6\x86\xd2\xcc u;a\x84\x95\x90\x12\xdb\v\x89N\xabp\x8b\xb5.\xefw\x9e\xbe\\{\xf8\x99\xa1Y\xe5\x1f\xbd$h=\x1e\x04N\xc1\x1a\\\xaa)\xa8c\x8a\xf3\x0f\xafsЛ\\\x81g\xabf5\x98\x97%H\xaf{\fC[ƻs\f\x9d3\xe1K\xbd\x8e\xce\xe1T\xaf\x18\xc1\xaa\x94\xde\x03\x00\x00\xff\xff\xfc\xee\x95g\x8d\x03\x00\x00",
		hash:  "e5451723f5b66ce13c8bf4d3804aa8c178a63b67b7caa5aa11b5731a35a5cb01",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755049, 0),
		size:  909,
	},
	"searchresults/type/ablock.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xbcU\xddj\xdb0\x14\xbeV\x9e\xe2\xcc\xf4rN\x18\xbd+\x8e\xa1n\v\x19ll\x8c\xbd\x80j\x1d\xc7\"\xb2d$\xa5[\x10y\xf7\x1d\xcb\xce\xf0\x16\xbbM7\xa8\xaf\x92\xf3}\xc7\xfe\xce\u007f\b\x02+\xa9\x11\x12\xfe\xa8L\xb9K\x8e\xc7\x05\v\xc1c\xd3*\xee\xc9\\#\x17h\xa39{\x97\xa6P\x18q\x804\xcd\x17\f2\x87\xa5\x97F\x83\x14\xeb\x04\u007f\xb6\xcaXb\xe6\v\x18\x9eL\xc8'(\x15wn\x9dX\xf3c\x84\xfc\x8d\x96F\xed\x1b\xed\x92\x1c\xfe\xa0DZ\xfd!\xbf\x15\x8d\xd4Pt\xfa\xb2\x15\xfd?'yR\x8f\xe7\xf6\x1e{$\xc9\xd3X\x8f\xdby\xb0'\x88\xfc\x931\xbb}\v\x1b\xee\xea\x9blE\x86\x17=BX\xf6N\x9d\xcf\xf1\xf8\xbc\x13\xa1\xf6?\x15\x16\xbc\xdc\xc17\xacТ.\xf1\x95J;\xe7\u07feo$x\x83r[\xfb\xcb5nb#.\xef\x8b\xde\xf1\r\x14~\xb5\xf8$\xcd\xde\xc1\xa8\xff.\xd4\xfb,!\x92x\x9c\x9a\x8a\x97\xde4\xa9Cn\xcb:UR\xef\x12\xf0\x87\x16קq\x1cE\xde\xc9\x19\nu*\x11\u007fA\xc9?f\x88\x90\xe9\x99!`zв\xfa:\u007f\xd0\xdeJtpg\xb4\xe7\xb4R\x04P\xd2F\xa9\xa3}\xd1p\xa5F\x11}F\xe7\xf8\x16\xef\xcc^S=axA\xb6\xeay4\xea\xd7\x13\v!\x04\xcb\xf5\x16\xe1J\xbe\x87+T\b7kX\xde\x16\xf7\xd2\xd1\xc6:О:\x11\x19\x1b\xf6B\xcc\xf4\x90\xe2hH\xce\x03`sk\x82\xb1\xc9^a\x9d]Ę\x0f\xf0\x9d*6\xd3\x18\x03/\x84N\xea\xb2#\xce5.\x11W\xd3_\x9a\xd7\xc4\xfa\xa7\xfbB\x97uԗ\xcfS\xd4\xf3\xc5\xd7h?\xeaʼZ\xd4\xe5-\xc2蜠\x16\xa3\xc2D\x16m\u007f\xba!\xec\xf4\x83\x8a\xde\x1f\x93|\xb83\x0fZ\x8cn\xcd\xf8\"\xb9\xd2\xcaֻdx\xe3\x18\xf2\xc6(wv\xc2*c|\u007f\xc2\x06%\xbf\x02\x00\x00\xff\xff)\a>\xed\xf5\x06\x00\x00",
		hash:  "d885b10188970b8c7bd9093ee5516b9fbbde629f218f31d3bb61962936399f9b",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755050, 0),
		size:  1781,
	},
	"searchresults/type/chainhead.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xecV\xc1n\xdb0\f='_\xc1\x19\xc1\xb0\x01s\xb3]]\xd5\xc0\x96\x04h\x80\xdd\xfa\x05\x8a\xcd\xd4Bdɓ\x94\xb5\x81\xe1\u007f\x1feˉ\xeb\xa6M\xdb\xf3|Hl\x92z|\xa4HJu\x9d\xe3V(\x84(+\xb8P\x05\xf2<j\x9a餮\x1d\x96\x95\xe4\x8e4^\x88\xa6\x15\xb3Oq\f\xbft~\x808N\xa7\x13`\x163'\xb4\x02\x91\xdfD\xf8XIm\xc82\x9dBxX.\xfeB&\xb9\xb57\x91\xd1\x0f\x03\xcdX\x9bi\xb9/\x95\x1dYԵ\xe1\xea\x1ea&\xbe\xc1\f%Br\x03WD\xc4\xeb&\x13b9\xdby\x11\xfe!\v\xf8>T\x88-\xccv'\xc1\x84ْK\x99.\xfa(\x13`\xbce\xbd\xe5\x99\xd3el\x91\x9b\xac\x88\xa5P\xbb\bܡB\x8ag#u\xb6\x8bR\xf2B\xae\xaf\x16Z9T\xee\xea\x96V7\r\x9b\xf3\x94\xcd;\xd0#e\xf2\x03\xac\xf8\xd1y\x81\xe0r\xbdL\xbe\x04\x88\xb5\xaa\xf6\xaei\xbe\xf6\v\x99\xad\xb8\x02\xeb\x0e\x92\xdcm\xa5\xe6.1\xe2\xbepQ\n#\xa7\xbfQݻ\xa2i`\xa5\x9c\x11h\t\x81\x96\x12\x03\xf26p_\xd7(-\x86\xb0;\x110\xc77\x949\x1fk\b\xb2\x15\x8c2\xdd\x1a3\xb7\xa1\xbd=\xa3 \x8d9'&y\x9ezJ\a\xb8\xe5\xb6H\u061c\xbe_\xb2\xbb\x9cp\x0ft\xccwHV\x97\xe8\xf3\xb0$7\xeff\xfb\xe8\xd0(.a\xbd\xb4\xaf\xf3\x9dN\xc2\xc3\xf6\xf2\xf8\xd1U%\x84\xb2\\/}\xf9=\xd9)\xc2'd\xf0\xddrZ\xd1V\xbb\x14]\x9b\xf8(c\f4b\x91\xb7\x11\xaf\x97>T)R\x18yB\x95\x0f\xc0\xd8|/\xcf\x13\xfeX\x8azXZ\x1c\x02x!'O\xfa\xf6\x92Ak\xd4\xd6\xf6)ଃ\x8f\xed\xbe,\xb9\xdf\xe6\xe0\x0f\xee:\x81o\xc84\xb4\xcc]\xa1\x1f\u09d4\xc7>\xa1\x1a\xb8\xecpc`\x9e\x02ͧ\x83C\x9b\x8c;(\xfc\xf7\x8d\xf4v\xb8\x9eg[\xe0/\xa0z\xdd{0W\v\x82\xb5\x0e\x9e\xe1\xad\x16^\xfe\x16\xa8n\x00\xd0ۇw\xc2\xf7z\xd4O\x9f\\X\x1a\xf7\x87Di\x85\xd7QJ\xc9\xef\x03O>ci\xab\xeb翃\r\xbb\x159\xfe߰ס\ns9%\xe7ɾ\xbd\x1c^o\xdb\xf7\x8d\b\x12\x9f?\rH\xe1\x0f\x90ѹ\xa3\xf2\x11\xcd {R\x9flNG\xbe\x9f\xa5\xfd\v\xf1\xeen\x10i\xb8\\\xachܝ.\x18\xc3k\x88͌\xa8\x9c\x8d\x82\x9b\xa1\xcai-\xed\xb3{\xcbVk\xd7\xdd[\x02\x95\u007f\x01\x00\x00\xff\xff\xbf\xbd}\xb8\xed\b\x00\x00",
		hash:  "9cae6493220b9658790e7c71938d38470ca1f300364642f144844d5511d8de55",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755050, 0),
		size:  2285,
	},
	"searchresults/type/dblock.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xd4W_o\xda>\x14}N?\x85\u007fQ\x1f\u007f\x01M}\xabB\xa4\x15\x8a\xa8\xa6IӶ/`b\xd3Xubd\x9bn(\xe2\xbb\xef\xfa\x0f!@\b\x84\xb6\x8c\xf1\x84|O\xec\xe3{\xae\x8f}˒\xd0\x19+(\nɔ\x8b\xf4%\\\xadn\x82\xb2\xd44\x9fs\xaca8\xa3\x98Pi\x87\xe3\xff\xa2\b=\b\xb2DQ\x94\xdc\x04(V4\xd5L\x14\x88\x91AH\u007fϹ\x90\x80Ln\x90\xffń\xbd\xa2\x94c\xa5\x06\xa1\x14\xbfj\x91\xddh*\xf8\"/T\x98\xa0-\x88\x85e\x9f\x92\x11\x93\xb0\x92\x90K\xf4`8\xc6}\x18\xdb\aj<\xe5t\u007f\xdcŦ@\xbb9\xe6\xe2\xf2p\xd0\x01H\xf2\x85.\xbf~\xbf\x8f\xfb\xf0\xf7(\xb6,{\x16\xbeZ\xb5\xe3!*\xdfH\xcb\nґ\xdbĊ\xda3\x9f^\x84\xe3x\xc19\x9a`\x95\x9dN\xd1|b\xbe\xb8\x00\xbb\x9f,\xa7?4\xce\xe7\x9d\x138\x162\x87SB\xaa\x19.\xa1\xb79\x02hB\xd9s\xa6;\x13\x1e=\xb8\x0f/\xc0\U000db92fL,\x14\xda9\xbd'rn\x05\x98_5\xbf\xad}to\xbe\xc3֊f\x18\x96\xcb#E\xb1L\xb3\x88\xb3\xe2%Dz9\xa7\x83\xb5\xc7\xd5\x12bf\xa9N*N\xe2\xa9D\xfd\x0ekW\x85m\xd7ߞvS\xc1\xed\x9b=S\b\x884\x9b\x1a\x04\x9a\x9d0\xce\xee\\\xf5(4\x14\x85\xc6`\xfb\x04\xb1\xc2\xc9\x02n\x9ec\xce\xeb\xee`\x86\x87bQ@\xb58\x8c\x8a\xfb\x0e\x04\x0e|\xd7\xe4\xd5va\xab\x81O\xbe\x1d\b\xf7\xa9\x04\x87\x1c9\b\x1a\x8b+0\xe3pRA\xc5\x03\xf5\xe3\x11\x9fI\xbe\xde\xd1!\\sJ\xdb\x176%\"\xdbW>Z|xS|\x96\xa5%i\xaf\t鋯\x1b\xe13\n\xe0d}\xde\xed\xd2lQl\x17\xfaXh\xf0\x88\xa1\xa4\x84\xe96\x05\xdfѣ\xdad\xdd\xc5\x1e\u0557\xa6\x1b\x81\xed^\xdcVN\x94\xf9C\x0e\xfcU\xeb=6\x99d\xe4_\x94z\xb6Q\xda\xef\xe2/\xaa\\\x96\x12\x17\xcf\x14ݲ\xff\xd1-\x05\xbd\xef\a\xa8\xf7\xe8\f\xbbv\xf7\x04\xc15\xf8\xb3;\xe4W\xe9ϴ\x12\xd5d\xb1W\u007f\x15td\x8aڨ\x1e{\v{\x98gq\xec\x19|&\x89a\x06\xb7?z\x1a\xbd1e\xa9\x99\xc6t\x88U\xd6\xfc\xdb\xc1\xce\xff4:/}\xad\x15&4\xe6\xc8T\x11\xa3\xea\xa4\fzFΑ\xddk\xe6\xc3.\xda\x00ZgZ\x90\xed\x17_܇N\x17\xfa\xe5`\xfd\a^Q\xaeqN|O\xfdX\x90Z_]\xef\xbeU*\xd9\\\xab\xd0\xcfX\x0fi!\xb8\xdak\xd7gBh\u05ee{&\u007f\x02\x00\x00\xff\xffq\xd0\xe5\xf7\xe1\x0f\x00\x00",
		hash:  "c0b6677fc67ef5062e5f8d7fedd45e6545285789d84f217496e9902761af4400",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755050, 0),
		size:  4065,
	},
	"searchresults/type/eblock.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xbcV\xddN\xe3:\x10\xbeN\x9fbN\xc4\xe5I\xab#\xeezB\xa4\x03\xe5\b\xb4BZ-O\xe0&Sb\xe1\xda]\xdb\x01\xaa\xa8\xef\xbe㟶)\x84\xd2v\x17r\xd5Όǟg>\u007f\x9e\xb6\xadp\xc6%B\x8aS\xa1\xca\xc7t\xb5\x1a$mkq\xbe\x10̒\xb9FV\xa1\xf6\xe6\xfc\xaf,\x83KU-!ˊA\x02\xb9\xc1\xd2r%\x81W\x17)\xbe,\x84\xd2\x14Y\f ~yş\xa0\x14̘\x8bT\xab\xe7\x8e絷T\xa2\x99K\x93\x16\xb0\x13\xe2\xc3\xea\u007f\x8aki\xf5\x12.\x1d\xbe|D\xff\xdf\x06Y6\x15\xf8\xd6\x1e|S\x82\xdc\xef\v~\xfd\xbe3\x04T\xc57\\\xde\xfd\x18\xe7#\xfa\xf9al\xdb\x0e}\xf8j\xb5?\x9e\xbc\xfa7a\xfd\xdf\b\x017\xccԇCsK܊/@wU3.o'\ab˙\xe7ь\x95V\xcd3\x83L\x97u&\xb8|L\xc1.\x17H$q\xe9\x1c\x1dSw\x8e\x1b\xcf\xcba\xdcÝ\x86\x15\x9f~\"OA\xb8A\xfeP\xdb\xc3K\x1e\xa1N.\xc3\xc2/\xa8\xfcw\x8dO\\5\x06:7\xe7@\xbc{\x03ܷ\xc9\xedI\xeeMc\xb7\xf6\xa3\xf6E}\xe9\x14\xc4e\xda\xdc\x14\xea\xdeT\xc3\xe8\x88\xfd7\xe4\xf7\xfb\xef\xa6ݲ|\xff\x81Ol\x04y\xfaE\x85\x1c\xfdJ\x94\xd7\xe7^\xc68\x1a\xb8R\xd2\x12k\xb1\x02.CoHJ\xe7L\x88Nm|\xe3\xaeT#\x89.\x10\x17\xe6\xa3\x10E\x1axޣ\x94m\xab\x99|@8\xe3\u007f\xc3\x19\n\x84\xf1\x05\f\xe3Ҟ:\xb4-\x9f\x01\xfe\xf4\xa1C_\xc6\xf4\x8eˆ4\xff\x8e\xe9\xc7 \xf9\xd0/\xb4\xbeӱ\xc5ސ~\xa2\xf4\xee\x80:\xf4\xce\xf9Cu\xb4\xe1\xab\xdaܶ(\fv*\x97$G\xd4,y\xaf`I\xd2[\xaa\xc4٫\xf8<\xeey\x06b\xdc\xc7W\xd4%J\xd7\xe5[\xbf\x12\xef\xe9*%\x1d\xf5\xa3\xda\v\xf6Ţ\x96L\xc0\xed\xc4\xec\x87;H\xe2\x977b\xf3'\x14\x19\"\xd3o'\x8e\xe4\x1e-奌\xe0\xe6\x94m\xa4\xef\x95\xe0a@q\x87\xcb0n\x9fq\xff\x8c\x9c\x05z\bN\xf7iw\a\x94U'Y>r\x18\xfa\x80\x9eV\x9auZ\xf7J\x92\x1a\x10\xb6?%ϹY0\xd99p\x19\xd2g\xa6\x99ϙ\xebn\xdc\x0f\xee\x83aL\xca]D\xf9\xb9\xaf\xd53\xfc'\xc4Vh֪L\x13\xa1nd\xe9&\xc2p\xb5B\x12W;\xb7\xdf\xe9\xb0\x1c\xdfS0v)\x88\u007f\x1574u.\xc7RI\xfc7-\b\t\xac\xab\xf3\x1aeD߇\xf4x\x80G6\xf1$Y\x90Վ*\xbc\xb6\xf8\xd54\x11;\xa2\xaf\u007f\x10\xf40`\x17q\xf6\xbe&Nn\xe7\xef\xee\x94nJ\xcd\x17֬5\xbb\xeb\xb2J\t\xf3f\xac\x9f)e\x83\xc6G$\xbf\x02\x00\x00\xff\xff+Q\x86\xa5\t\f\x00\x00",
		hash:  "909c9767e9d7b7300eb04417ff26ad99ea315b9e7bad9f5c9f66c0ea34e62167",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755050, 0),
		size:  3081,
	},
	"searchresults/type/ecblock.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xbcUMo\x1a1\x10=ï\x98\xaer\xec\x06U\xb9E\xcb\x1e\xa0\xa89\xf4п`\xd6\x03k\xc5\xd8\xc8v\xd2\"\x8b\xff\xde\xf1\a\xd56\xec.\tH\xe1\xb4\xcc{\x1e=?χ\xf7\x1c7B!\x14ج\xa5n\x9e\x8b\xe3q:\xf1\xde\xe1n/\x99\xa3x\x8b\x8c\xa3\x89\xe1\xeaKY\xc2B\xf3\x03\x94e=\x9d@e\xb1qB+\x10|^\xe0\x9f\xbdԆ\x98\xf5\x14\xf2\xaf\xe2\xe2\x15\x1aɬ\x9d\x17F\xff\xee o\xd1F˗\x9d\xb2E\r\xffQ\"\xad\xfdV\xaf\x943\aX\x1a\xe4\xc2\xc1\"Ȭf\x14>\xe7:\xb6\x96x\x1eOؚ\x94\xf7c\t7\xc3`\"\xf0\xfa\x89\xd9\xf6\xb1\x9a\xd1\xd7E\xaa\xf7\xf7\xabe\xd4z\xff\x03]8x<\x8e\x9f$\xd4ܨ/>\xce\xf5\"\xe3S\x87\xaf\x90\xe7\x93$?\xa1ض\xeeF\xbd\xdf\x17)\xcd'\xe8\xfde\xf0U\xe8\x17\v\xe7E\xf9\xceK\x8c\x12\"\x89Ŏڰ\xc6\xe9]i\x91\x99\xa6-\xa5P\xcf\x05\xb8\xc3\x1e\xe7\xffzuЏ 2\xfd;\xbd\"\xbb\xa0\xebJ\xdb\b\xe9o+\x02\x06z\xb1j\x1fbC\v\xb4\xb0\xd4\xca1\x1a?\x1c\x84\xea1\x94&̎I\x19\xae\xf9\x13\xd5\xd6\xd1M \x1f\xadf\t\xa2A\xf0\xd035\xf2(\x88>f\x03c\xa0\xb8n8xo\x98\xda\"܉\xafp\x87\x94\xf7q\x0e]\xe7\xb3(\x1a\x92\xc3\x19\xc4&\x1e=\xd1s\u007f\xddV\x8cɲ\x0f4\xfc\xe5\xd2\n\x19Ca\xf5h\rU\xf4\x8eB\x01\x18q\x01\x15\x1fui\x18\xffP\xa9U3\xda.\xb4\xa3&\xa7\x0f\xaa\x97\xb4\xac\xea\xbc\xc7V\x8awvYw\xe3\xd9ƈ\xbd\xb3E\xd6х\x9c\xd6Ҟ\xadȍ\xd6.\xadȬ\xffo\x00\x00\x00\xff\xff\xfeґ\xadV\a\x00\x00",
		hash:  "5a92072cac8a8cafae3c30cd758ec30320b45160754be5716e3753b8da4fc2d6",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755050, 0),
		size:  1878,
	},
	"searchresults/type/ectransaction.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xc4TMn\xf3 \x10]'\xa7\x98\x8f\xbd?\xabۊx\xd1*j\x0f\xd0\v\x103\x91Q\tX0\xfd\x89\x90\xef\xde18\xaa\xabF\xa9\x17U\xcaʞ\xf7`\x86\xe1\xbdII\xe3\xde8\x04\x81-\x05\xe5\xa2j\xc9x'\x86a\xbdJ\x89\xf0\xd0[E\x8cv\xa84\x86\x1c\x96\xff\xaa\n\xee\xbc>BU5\xfc\x1b1o\x01\xa37\x02\xdf{\xeb\x03\x13\x19XIm^\xa1\xb5*ƍ\b\xfe-Ǿ\x04[o_\x0e.\x16\x00@v7\xcd\xd6Q8\xc2}@m\b\x9e>\v\x925\x83k8\xb3$\xa9\x9dŜ=\xa2\nmW\xe5\x808\xcf>\xed\xd9q\xfd\x97\x18\x99\x14~`\x14\x96^\xc0\x1a\xd7\xec6\xf0\xa8bw\xbb\xe0\xf0z\xd1\xe9\xcbkH\xe9\xff\x03Ҙ\x9d_\xf27\xd23%\\\xb9\x8fE\"\u007f\xd5A\xa9\xb2\xd4\xf6\xfc\x90\xfePM\x8a\xb3\xc6=\v\xa0c\x8fl\x81\xb1<єN\xe7ZK\xbbe\xad\x9ak\xb4\x9c\xe1\v\xe2fttG\xb1b\xcd^\xccF\x9d>d=y\xb9\x99\\\xbeuz\xe6\xf4\xf9<\x88m0=Eq\x92\xd1\x1c#\xefm\xfc6A\xf6\xdeS\x99 )\xa1\xd3\xc3\xf0\x11\x00\x00\xff\xff\xaeĢ\x15{\x04\x00\x00",
		hash:  "7d590b3a71bd677ee47f433a0177e1d1de7eb5fee91da29a4e0a5e88fc34dbcd",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755050, 0),
		size:  1147,
	},
	"searchresults/type/entry.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xe4U\xcdn\xdb0\f>gO\xc1\x19\xc5n\xae\xb1\xab\xab\x18ؒ\x00\r\xb0[\x9f@\xb1\x99Z\x98,\x19\x92\xb2\xd6\x10\xfc\xee\xa3$\a\xf1f\f1v\xad\x0e\x01\u007f\xe2\xef\xa3H\x8a\xf4\xbe\xc1\xb3P\b\x19*g\x86l\x1c?m\xbcw\xd8\xf5\x92;\xb2\xb6\xc8\x1b4\xd1\xcc>\xe79|\xd7\xcd\x00y^\x91j\xb1vB+\x10\xcd6\xc3\xf7^jC\u007f$ǆ5\xe2\x17Ԓ[\xbb͌~\x8b\xb6?\x8c\xb5\x96\x97N\xd9\xe4ذ\xf6ku\b\xe4\xb0玳\x82\xd4dw\xfc$1ɤ\x9c\x02s\xe0\xb2\xc8M\xdd\xe6\xd1;A\x04\xbf\xb9\x8a$7\x13\xe03\xb7m\xc9\n\xd2\xe7>\xef\x1f\x83c\x1c\xe7\x1e\x92Ϳ\xc0v-\x17\n\x8e\xfb\x05\x12\xe31\xa23\xaf\x9d\xee\xf2)0)\xd4\xcf\f\xdc\xd0#\xdd4|\x19r\x98\x05\u0588s\xdc\ab^\xad%?\xbc;4\x8aK\xe2\xb7˻\\\xe5\r\xbbț\x02t\xbc\a\xc3\xd5+\xc2\xc3q\x0f\xe5\x16\x1e\t\x88  \x94r\xf6\xb7p\x98\x14\xa9\x8a!i9N\x84\xb9\x88Q?\xa4\x80\xa5\xa8\xe0o|T\xcd\x1c\x8e\x15\xb3\x18V\xe7V+G\xbc%;\x19(\xaa\xe5\xfd\xe0\xcea\xb6\xe7j\x16}\x9d\xf0r{\xe9:N\r}%\x80\x97d(\x81Q\xeemǥ\xac^Z\xfd\x06ߤdE\xd2CY\xee\x13\xc6@\x81\x9e\xc2\xe0Ж\x10\xea\x9a(~\xa0zu\xd4V\xeb!\xae\xb1\xc5>\x9d!\xa5\xf6\\\x8fs\xd8\x11\x94u\x101\x0e\xbb \xaf\xf9\xbc\b\xb9\xabH\xfa\xef,\x87G\x99\x81u\x83\xa4fo\x84\xa5\xa91\x94J+|\xca*J\xec\xf5\x82\xe5\x17\xecl\xff\xb4\xfc\x9d\x15\xe3Y4\xf8\x91\x8bњ\xfb\u05fd\x05\xb5\xbe\xbcw\x1e$\x89\xa1\x88\xd3\xc8-n3\x97\x154\xb1\xe38\x9f\x04\xc2K\x13\xbf\x9av\xc1\x81\x9e\xffm\x1f̷\x86\xad\x8d\xe8\x9d]l\x13\xa7\xb5\\Z\xcfZ\xbb\xb4c\xbc\xa7\x912\x8e\xbf\x03\x00\x00\xff\xff\x1a\x98\xf9\x91\x95\x06\x00\x00",
		hash:  "a9b9fe61fc7a2385e955d54519ae2bb7a9da16621e812e7030e44ee0fc63100b",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755050, 0),
		size:  1685,
	},
	"searchresults/type/entryack.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xbc\x94MN\xeb0\x10\xc7\xd7\xe9)\xfc\xbcO\xa3\xb7En\x16@%\xf6p\x81i\xec*V\x1d;\xb2\xa7@e\xe5\xeeL퀂hJ\x16\x80W\xb1\u007f\x93\xf9\xd2\xfc'F\xa9\xf6\xda*ƕE\u007f\x82\xe6\xc0\x87aUĈ\xaa\xeb\r \x81V\x81T>=\x8b\u007fe\xc9n\x9d<\xb1\xb2\xac\xe9\x1aT\x83\xdaY\xa6冫\xd7\xde8O\x86\x04\n!\xf53k\f\x84\xb0\xe1\u07bd\xa4\xb7O\x8f\x8d3\xc7Ά\f\n\xd1\xfe\xaf\xb7\xe7\xf8\xecɃ\r\x90\xbd>\"\xe01\x88\x8a\xe0\x8a]8\x02ag\xd4eV\b\xdcQ\x9e3\x90\xa8\x9fC\xc4dN\xe6\x01B{#*\xba^3\x15\x90\xca\xdfSҮ+\x83\x02ߴ\xa5\xd1\xf6\xc0\x19\x9ez\xb5ɍ\xe5u\x8c\xeb\x0f\xaf\xc3 *\xa8\xaf\xb9&6\x93\xe1\xb4\xfeoMF;Y߹\xae\xd38\xb6\xf4JQ\x17~]b\x97\x0eU\x98\xc3\xdc\x03\xc2:\x87\xa2\xa9Y\x12fA>?ݑ<o\xbfݐ\x14\xe5\x8f\xfbq\x1e\x9e\xf9\xe1'\x98es\x9eߊ$\x99\xf4:~\x88j\x94t=\x8a}k\xe5D\xf0ӵ\x10\x1a\xaf{\f\xfc\xbd\xa2)C\xe7L\xf8\xb2H\xf6\xcea^$1*+\x87\xe1-\x00\x00\xff\xff\xe2\x95\b\x0e}\x04\x00\x00",
		hash:  "bdd438304a91ea0bbb47d4b261e0ec89fa0888139f2b8d47ff1353b21492e7d0",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755050, 0),
		size:  1149,
	},
	"searchresults/type/factoidack.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\x9c\x92\xc1n\xc20\f\x86\xcf\xe5)\xb2\u07bbj\xd7)\xe40\xb1I\x9c\xc7\v\x98$\x88\x88\x90T\x89\xd9@Q\xdf}n´\"ʘƉ\xfas\xdd?Η\x92\xd2\x1b\xe34\xab7 \xd1\x1b\x05rW\xf7\xfd\xacJ\t\xf5\xbe\xb3\x80\x84\xb6\x1a\x94\x0e\xb9\xcc\x1f\x9a\x86\xbdxubM#\xe81j\x89\xc6;fԼ\xd6\xc7\xce\xfa@\x8d\x04*\xae\xcc\a\x93\x16b\x9c\xd7\xc1\u007f\xe6\xdaEQz{ػX@ŷO\xe2\xad$`\xab\x00.B\x99\xfb\x8e\x80\x87\xc8[\xc236\xf1\xe3\bk\xab\xa7Y\xc5qMIo@\xa2\xe1\x1a\xe5,\xa8\xc48\xc3r\xf1\xcc[\xaaMι\f\xa3\x04\x87\xbc\x8b\xbc\xcd}\x135\x04\xb9m\xacq\xbb\x9a\xe1\xa9Ӆ\xe0\xcf\xf4Z\xa4\xf4\xb8:.\x17}\xcf[\x10\xf7?D\x1d\xe1N\xc7o\a\xfb^\xe8_\xcfC\xe9\xca+C\xbe\xffg\xab\bݾ\f\x82\xe5\x1a\x87\x90-I\x92\r:\xff\xe1\xedY2q\xd6\xefթ\x91\x82cQ\xa3\f\xa6\xc38\x98:\x8c\x1d#\xf4\xde\xc6+\xb37\xdec1;%\xedT\xdf\u007f\x05\x00\x00\xff\xff}\xcag\xb0\x10\x03\x00\x00",
		hash:  "d3778c57993f99c57a010e02aaf2088588c9f9e59abce76a0219877b96a1faf4",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755050, 0),
		size:  784,
	},
	"searchresults/type/facttransaction.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xdcWmo\x9b0\x10\xfe\xdc\xfe\n\xcfۤM\x1a\xa5\xab\xba/)A\xca\x18\xd9\xfai\x1f\xb2?\xe0`'X%6\xb2\x8f5\x11\xca\u007f\xdfa e\xea\x1b\xec\xa5IKՄ\xd8\xe7\xbb\xf3\xf3\x9c\xce~ʒ\x8b\x85T\x82\xd0\x05K\x00\fS\x16\xbf\xa5Vt\xbb=>*K\x10\xab<c\x80\xf3\xa9`\\\x187\x1c\xbc\xf2<\xf2Y\xf3\r\xf1\xbc\x10\u007fZ\xe1\x96\x10\xc9\xc7T\xac\xf3L\x1b4ĉ\xa3\x80˟$ɘ\xb5cj\xf4\xb5\x1b\xfbm0\xd1Y\xb1R\xb6\x9e $H?\x86S\x8c\xaf%'?nr\t|\x1c?&w<\x01\xb0y&\\`+\x98IR\xcf\rл\xad\xdb5sL\xfd!\x8b\xda\xca\x10\v\x9bL\x8c\xe9\x9c%WK\xa3\v\xc5=\xccW\x9b\xd1\xeb\xb3\xd3ꏆ\x01\xf00\xf0\xab\x8f\u074b\x0f\xa6\x87\xeb\xc7L\x9c\x15'\x18\xce\xe6L\x8d\xe9\x19\r\xef\xb7|\xdc\xd7\rTa?\xe36\x81\x06\x81k\xc9!\x1d}:}{\x01b\r\x1e\xcb\xe4R\x8d\x12\xa1@\x98\v:\xc4ez\x1e^\xaa\xbc\x00\x8b\x94\x9e\x0fXX\x96X\rKA\xde\xc8\x0f\xf8\xaf\xc8hLN\xbe\n\xa8}aI\x92!O\xc0\\\xc1T\x05\xafW^S7\x99TW\x94\xc0&\xc7\xedN'4,\xcb\t\xe7FX;\x9dD\xda\x18\xac\xf0*p\x15\xb4\x19?\x99\x81\x91j\xb9\xdd\x06>\vI07\xc4\x1f\xb4!\xa1\xf8\x90\xc4]u\xfd\x11y\xb7(\x1b\xca\xd8\xf7\x02\xfe\x962]@\xcbY\xe3\xed\xc9H\xc3\xd0\xfbc\xed>\x10\xe2\xe8\u007f\xc1\x10G\x1d\x18\xe2\xa8\x0f\f\x87T\xbbhڳM\xf5pڣ\x19\xf7\xee\xc5=\xf3\xef\x9cZ\xe4\xf2\vy\ak\xc9ߏ\x8e\xff\x15D\xfd\x13)ˊ\xee\x99\\~c6\xedC\xd6\xe1\xe3Y\xedd\u007fP>s\x1c5\xb0\x8c\xb8\xf3r?\x10v\x88\x9c\xac\xf0.\x05mg:q\xa9\xb5'9i\xae\u007f\xf6y\x03\xdd^b\xeb.\u007f\x90\x88\xef\x0e\xa0\x17\x02y\x1c\x1d2\xdaq\xf4\"\x90\xeet\xe3\x190(\xec\x9e\xfaq\x1d\xfci\xda1N? \x18w\x17\x96J\xd9\xfa(m\x9d\xeem^\x02\xbf\x91\xc6a#\x9ac\xc5;¹+\xafmbd\x0e\x96\xb6{\xea\u0381F\x1dxK\x90/\xb4\x86Z\x907\x17\xb2_\x01\x00\x00\xff\xffy\xac\x81\xde\xcc\x0f\x00\x00",
		hash:  "87249c7c9dd50e7b313ec61605b35b41978f0a5946f208ad8db955ddb3bd0b82",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755050, 0),
		size:  4044,
	},
	"searchresults/type/fblock.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xff\xd4X\xdbr\xda0\x10}\x0e_\xa1\xba\xe9[\riҾ\x10\xe3\x19B\x9d\xa6\xd3f\xdaI\xfb\x03\xc2\x12X\x13!1Ғ\xcbx\xf8\xf7\xaedӘb\xc0\xb9\xd1ļ\x18\xedJ:\xbbg/\x92\xf3\x9c\xf1\x91P\x9c\x04\xa3\xa1\xd4\xe9e0\x9f\xb7\xf6\xf2\x1c\xf8d*)\xe0p\xc6)\xe3\xc6\x0fGo\u0090\x9chvK\xc20n\xed\x91\xc8\xf2\x14\x84VD\xb0^\xc0o\xa6R\x1bԌ[\xa4|\"&\xaeH*\xa9\xb5\xbd\xc0\xe8\xeb\x8a\xe4_i\xaa\xe5l\xa2l\x10\x93%\x15\xaf\x96}\x88Oi\nZ0r\xe2\x10F\x1d\x1cYU\x03:\x94|u\xbc\x90\r\x11t\xbd\xac\x90\x9b\xf5\xc2B\x81\xc5\xde\xee\xf3\x8bn\xd4\xc1?[\xb5\xf3\xbc\xed&\x9c_\xcc\xe7\x9b'\xa0\xd4<\x12\xd9\x19\x17\xe3\f\x9a\x03\xfb|R\xcc\xd8\x01\xb4\xe4&ͨ\x1asr\x81\xb1\xd4\x1c\xa1\x9b\xe6f\xec\x00\xe1Oï\x84\x9eY\xb2\x14c\r\xa1nT\xf0J\xd4\xe7\xc6ȭ=\t-\xa7&\xcdB)\xd4e@\xe0v\xca{\x8b\xa4sF;$\xdfx\x192t\xcb\xe6\x0ft\vJ\xeaS\x01\x05\xf5\xf9\x13eG\xf1oC\x95\xa5>\xd5-\x19h\x05\x14\v\x06#B-;\r\xeb\xc1\x84J\xe9l\xf9\xce\xd5\x18\xb2\xf9\x9c$\n\x8c\xe06\xea\x14\"\xccݣ\x9a\x1c\xcfs\xe3\xa3d_\xbc'\xfb\\r\xd2\xed\x91vuW\xac>\xa4>\xe3\xbd\u007fK\xc7\xfa\x81`\x8d_\xb6\x15\x01\x17.\xc4\u00adDV\x864\xbd\x1c\x1b=S,\xc4ʤM\xf7\xed\xe1\x81\xfb\x05\xb1c\xdd;\xff\xeeec\f6\nB\x82\x9b\xd8)U\xbd\xe0\x10+\xe0\xdag{\xb4m\xa8\x81uۖ\xd6^\v\x06Y\xf7\xd3\xc1\xbbc\xe07\x10R)ƪ\x9br\x05\xdc\x1c\a\rW\xcb>\xc6_\xd5t\x06\xc84\xbe6\x9a\xb3D:\xc6\x12r\xee\xb8o\u007f\xe1P,UG\xfaC\xf3\xec\xb4\xefr\xacϘ\xe1֞\xf6\a\xda\x18\xec]n_\xb7_9\xde\xfe\x85\xc1\xaa\xc6E\xfe\x91hhH\xa7\xa9)\\\xb1\x86p\xb7W\x96\x1a\x8aV\x88\xb9\a/?f\xf0\bb\xf4\f\xaa̔\x8b\xed\x82\x1a\xdcy\xd7\xdcl\xb2<\x19<\x83\xedɠb{2hb\xfb\u007f\r˵m\xe2)\x9aS㞽\x15f\xa5u\x903j\xb3n\xebq\xa6?Q\xafG\t\xdc!\xf3M\x1fIv\x00\x9f\xb1\xe5?\xa1W5PI|mމC\xf3\xbc\xc2c\u007f\x82\xed\x18\x16\t\xd2\xf6X\x16mbq\f\xb1\xaf\u0081\x8b3SQN^\x86'\xff\x96\xb6\xd7\xe5\xcad𢼘\f\x9e݃\xf7?\xc0\xd74\x81\xa8\x83Wo\xbc\xc0\xef-^\xf0x^\xdc\xe4\xe3\xf2\x92\x9f(V\xb9\xe8W?\a\xd8Ԉ)ؠ\\\xb1*\x02\x8d\xc7ؕ\xef\a#\xad\xa1\xf8~P\"\xf9\x13\x00\x00\xff\xff\x8a\x8c\\/r\x10\x00\x00",
		hash:  "0e4acdd241874ffef73fb1ea1efd38e23daf97e38822a335bcf44216e4c1cffd",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755050, 0),
		size:  4210,
	},
	"searchresults/type/notfound.html": {
		data:  "\x1f\x8b\b\x00\x00\tn\x88\x02\xffdRA\x8e\xdb0\f<o_\xc1\xe6\x9c\xc4\xe8\xb5\xd0\xfaP\xa0\x05\xf6\xd2K\xfb\x01ڢ#bmѐ\xe8d\r\xc3\u007f/-g\xb7)z\xb3f\xa8ь9\xcb\xe2\xa9\xe3Hp\x88\xa2\x9dL\xd1\x1f\xd6\xf5\xd3Ӳ(\rc\x8fjD \xf4\x94\n\xec>\x9fN\xf0M\xfc\f\xa7Sm\xc7L\xad\xb2D`\xff|\xa0\xb7\xb1\x97d\x83F<9\xcfWh{\xcc\xf9\xf9\x90\xe4V\xb0\u007f\xc0V\xfai\x88y'\x9e\\\xf8R\xff\x14\x85\x1f\x9b\x01W\xd9i\x87\xc7\xfaE!\x13\r\x194P\"\b\x98\xa1!\x8a\x80\x11(%Ip\v\xdc\x13h\x9a9^@\x05Z1\xdfd\xbe5\xe0v\x17S\x1b\xce\xf0;\x10\xb0E\xb2\xcb\xdb\xdc\x0e\x93\x87\x9b\xe9Yp(ɏ\xc0\x1d\xcc2A\xc4+_,\xbb\xdf\x045p\x86\x11/\x04\xcdl\xee\xb9}\xdd\x14\x10z\x8e\xaf\xc7\xeda\xe08N\xaa\xe5}{f\xd7\x06\xa54\x18S\x04>\xc0\x06\x13`#W:\xc3K\xb7;\xec\x90\xfb||\x1c\x1ap\xb6\x9cW*\x94y\xf0\x13m>:lU\x06\u007fO\xe0\xe5\x16\xc1\xe2\xeb{\xb0-\x05\xbdq\xe6b\x84ca<*6\x98\xe9\xec\xaa\xf1\xe3\x97\xfe\xfak\xf0+\xb8\xa6^\x96\U000faeaa\xa9߇\\e\x8b*[\xbc\u007f\xb8\xea\xbe\xe8\xfa^\x81\xef\xd1?\xd4\xe0\xb1,\xb9M<j\xfe\xafD\x9d\x88\xee%Z\x16\x8a~]\xff\x04\x00\x00\xff\xfft8\xff\xe9y\x02\x00\x00",
		hash:  "f6de428c64d09d9ace9dd83a38ecbddf09b543f5747e06f555262ebf01ea4d67",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1472755050, 0),
		size:  633,
	},
}

// NotFound is called when no asset is found.
// It defaults to http.NotFound but can be overwritten
var NotFound = http.NotFound

// ServeHTTP serves a request, attempting to reply with an embedded file.
func ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	f, ok := staticFiles[strings.TrimPrefix(req.URL.Path, "/")]
	if !ok {
		NotFound(rw, req)
		return
	}
	header := rw.Header()
	if f.hash != "" {
		if hash := req.Header.Get("If-None-Match"); hash == f.hash {
			rw.WriteHeader(http.StatusNotModified)
			return
		}
		header.Set("ETag", f.hash)
	}
	if !f.mtime.IsZero() {
		if t, err := time.Parse(http.TimeFormat, req.Header.Get("If-Modified-Since")); err == nil && f.mtime.Before(t.Add(1*time.Second)) {
			rw.WriteHeader(http.StatusNotModified)
			return
		}
		header.Set("Last-Modified", f.mtime.UTC().Format(http.TimeFormat))
	}
	header.Set("Content-Type", f.mime)

	// Check if the asset is compressed in the binary
	if f.size == 0 {
		header.Set("Content-Length", strconv.Itoa(len(f.data)))
		io.WriteString(rw, f.data)
	} else {
		if header.Get("Content-Encoding") == "" && strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
			header.Set("Content-Encoding", "gzip")
			header.Set("Content-Length", strconv.Itoa(len(f.data)))
			io.WriteString(rw, f.data)
		} else {
			header.Set("Content-Length", strconv.Itoa(f.size))
			reader, _ := gzip.NewReader(strings.NewReader(f.data))
			io.Copy(rw, reader)
			reader.Close()
		}
	}
}

// Server is simply ServeHTTP but wrapped in http.HandlerFunc so it can be passed into net/http functions directly.
var Server http.Handler = http.HandlerFunc(ServeHTTP)

// Open allows you to read an embedded file directly. It will return a decompressing Reader if the file is embedded in compressed format.
// You should close the Reader after you're done with it.
func Open(name string) (io.ReadCloser, error) {
	f, ok := staticFiles[name]
	if !ok {
		return nil, fmt.Errorf("Asset %s not found", name)
	}

	if f.size == 0 {
		return ioutil.NopCloser(strings.NewReader(f.data)), nil
	} else {
		return gzip.NewReader(strings.NewReader(f.data))
	}
}

// ModTime returns the modification time of the original file.
// Useful for caching purposes
// Returns zero time if the file is not in the bundle
func ModTime(file string) (t time.Time) {
	if f, ok := staticFiles[file]; ok {
		t = f.mtime
	}
	return
}

// Hash returns the hex-encoded SHA256 hash of the original file
// Used for the Etag, and useful for caching
// Returns an empty string if the file is not in the bundle
func Hash(file string) (s string) {
	if f, ok := staticFiles[file]; ok {
		s = f.hash
	}
	return
}

// Slower than Open as it must cycle through every element in map. Open all files that match glob.
func OpenGlob(name string) ([]io.ReadCloser, error) {
	readers := make([]io.ReadCloser, 0)
	for file := range staticFiles {
		matches, err := path.Match(name, file)
		if err != nil {
			continue
		}
		if matches {
			reader, err := Open(file)
			if err == nil && reader != nil {
				readers = append(readers, reader)
			}
		}
	}
	if len(readers) == 0 {
		return nil, fmt.Errorf("No assets found that match.")
	}
	return readers, nil
}
