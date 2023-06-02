// Package bus provides a system Bus connecting all main system parts.
package bus

import (
	"github.com/retroenv/retrogolib/arch/nes/cartridge"
)

// Bus contains all NES sub system components.
// Since many components access other components, this structure
// allows an easy access and reduces the import dependencies and
// initialization order issues.
type Bus struct {
	Cartridge   *cartridge.Cartridge // used by Mapper
	Controller1 Controller           // used by Memory
	Controller2 Controller           // used by Memory
	CPU         CPU                  // used by PPU
	Mapper      Mapper               // used by Memory and PPU
	Memory      Memory               // used by CPU
	NameTable   NameTable            // used by CPU and Mapper
	PPU         PPU                  // used by Memory
}
