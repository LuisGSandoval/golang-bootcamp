package retroclock

type placeholder [5]string

var zero placeholder = placeholder{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

var one placeholder = placeholder{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

var two = placeholder{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

var three = placeholder{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

var four = placeholder{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

var five = placeholder{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

var six = placeholder{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

var seven = placeholder{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

var eight = placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

var nine = placeholder{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

var colon = placeholder{
	"   ",
	" ░ ",
	"   ",
	" ░ ",
	"   ",
}
var dot = placeholder{
	"   ",
	"   ",
	"   ",
	"   ",
	" ░ ",
}

var digits = [...]placeholder{
	zero, one, two, three, four, five, six, seven, eight, nine,
}

var letterA = placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"█ █",
}
var letterL = placeholder{
	"█  ",
	"█  ",
	"█  ",
	"█  ",
	"███",
}
var letterR = placeholder{
	"██ ",
	"█ █",
	"██ ",
	"█ █",
	"█ █",
}

var letterM = placeholder{
	"█ █",
	"███",
	"█ █",
	"█ █",
	"█ █",
}
var admiSign = placeholder{
	" █ ",
	" █ ",
	" █ ",
	"   ",
	" █ ",
}
var empty = placeholder{
	"   ",
	"   ",
	"   ",
	"   ",
	"   ",
}

var alarmSign = [...]placeholder{
	empty, letterA, letterL, letterA, letterR, letterM, empty, admiSign, empty, empty,
}
