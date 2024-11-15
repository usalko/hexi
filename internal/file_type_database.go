package internal

import "regexp"

type FileType uint16

const (
	SHEBANG              FileType = 1   //Script or data to be passed to the program following the shebang (#!)
	CLARIS_WORKS         FileType = 2   //Claris Works word processing doc
	LOTUS_123_V1         FileType = 3   //Lotus 1-2-3 spreadsheet (v1) file
	LOTUS_123_V3         FileType = 4   //Lotus 1-2-3 spreadsheet (v3) file
	LOTUS_123_V4_V5      FileType = 5   //Lotus 1-2-3 spreadsheet (v4, v5) file
	LOTUS_123_V9         FileType = 6   //Lotus 1-2-3 spreadsheet (v9) file
	AMIGA_HUNK_EXE       FileType = 7   //Amiga Hunk executable file
	QUARK_EXPRESS        FileType = 8   //Quark Express document
	PASSWORD_GORILLA     FileType = 9   //Password Gorilla Password Database
	LIBPCAP              FileType = 10  //Libpcap File Format
	LIBPCAP_NS           FileType = 11  //Libpcap File Format (nanosecond-resolution)
	PCAPNPG              FileType = 12  //PCAP Next Generation Dump File Format
	RPM                  FileType = 13  //RedHat Package Manager (RPM) package
	SQLITE3              FileType = 14  //SQLite Database
	AMAZON_KINDLE_UP     FileType = 15  //Amazon Kindle Update Package
	DOOM_WAD             FileType = 16  //internal WAD (main resource file of Doom)
	ZERO                 FileType = 17  //IBM Storyboard bitmap file, Windows Program Information File, Mac Stuffit Self-Extracting Archive, IRIS OCR data file
	PALM_PILOT           FileType = 18  //PalmPilot Database/Document File
	PALM_DSK_CALENDAR    FileType = 19  //Palm Desktop Calendar Archive
	PALM_DSK_TODO        FileType = 20  //Palm Desktop To Do Archive
	PALM_DSK_CALENDAR2   FileType = 21  //Palm Desktop Calendar Archive
	TELEGRAM_DSK         FileType = 22  //Telegram Desktop File
	TELEGRAM_DSK_ENC     FileType = 23  //Telegram Desktop Encrypted File
	PALM_DSK_DATA        FileType = 24  //Palm Desktop Data File (Access format)
	ICON                 FileType = 25  //Computer icon encoded in ICO file format
	APPLE_ICON_FORMAT    FileType = 26  //Apple Icon Image format
	THREE_GPP            FileType = 27  //3rd Generation Partnership Project 3GPP and 3GPP2 multimedia files
	HEIC                 FileType = 28  //High Efficiency Image Container (HEIC)
	Z_LZW                FileType = 29  //compressed file (often tar zip) using Lempel-Ziv-Welch algorithm
	Z_LZH                FileType = 30  //Compressed file (often tar zip) using LZH algorithm
	LZH0                 FileType = 31  //Lempel Ziv Huffman archive file Method 0 (No compression)
	LZH5                 FileType = 32  //Lempel Ziv Huffman archive file Method 5 (8 KiB sliding window)
	AMI_BACK             FileType = 33  //AmiBack Amiga Backup data file
	AMI_BACK_IDX         FileType = 34  //AmiBack Amiga Backup index file
	BPLIST               FileType = 35  //Binary Property List file
	BZ2                  FileType = 36  //Compressed file using Bzip2 algorithm
	GIF                  FileType = 37  //Image file encoded in the Graphics Interchange Format (GIF)
	TIFF                 FileType = 38  //Tagged Image File Format (TIFF)
	BIG_TIFF             FileType = 39  //BigTIFF
	CANON_RAW_V2         FileType = 40  //Canon RAW Format Version 2 Canon's RAW format is based on TIFF
	KODAK_CIN            FileType = 41  //Kodak Cineon image
	RNC_V1               FileType = 42  //Compressed file using Rob Northen Compression (version 1) algorithm
	RNC_V2               FileType = 43  //Compressed file using Rob Northen Compression (version 2) algorithm
	NURU_IMAGE           FileType = 44  //nuru ASCII/ANSI image files
	NURU_PALETTE         FileType = 45  //nup	nuru ASCII/ANSI palette files
	SMPTE_DPX            FileType = 46  //SMPTE DPX image
	OPEN_EXR             FileType = 47  //OpenEXR image
	BPG                  FileType = 48  //Better Portable Graphics format
	JPEG_RAW             FileType = 49  //JPEG raw or in the JFIF or Exif file format
	JPEG_2000            FileType = 50  //JPEG 2000 format
	QUI                  FileType = 51  //The “Quite OK Image Format”
	IIF_ILBM             FileType = 52  //IFF Interleaved Bitmap Image
	IIF_VOICE            FileType = 53  //IFF 8-Bit Sampled Voice
	IIF_AMIGA_CB         FileType = 54  //Amiga Contiguous Bitmap
	IIF_ANI_BMP          FileType = 55  //IFF Animated Bitmap
	IIF_ANI_CEL          FileType = 56  //IFF CEL Animation
	IIF_FAX_IMG          FileType = 57  //IFF Facsimile Image
	IIF_FT               FileType = 58  //IFF Formatted Text
	IIF_MUS_SCORE_SIMPLE FileType = 59  //IFF Simple Musical Score
	IIF_MUS_SCORE        FileType = 60  //IFF Musical Score
	IIF_YUV_IMAGE        FileType = 61  //IFF YUV Image
	IIF_AMIGA_FVM        FileType = 62  //Amiga Fantavision Movie
	IIF_AIFF             FileType = 63  //Audio Interchange File Format
	LZ                   FileType = 64  //lzip compressed file
	CPIO                 FileType = 65  //cpio archive file
	DOS_MZ               FileType = 66  //DOS MZ executable and its descendants (including NE and PE)
	SMART_SNIFF          FileType = 67  //SmartSniff Packets File
	DOS_ZM               FileType = 68  //DOS ZM executable and its descendants (rare)
	ZIP                  FileType = 69  //zip file format and formats based on it, such as EPUB, JAR, ODF, OOXML
	RAR_V1               FileType = 70  //Roshal ARchive compressed archive v1.50 onwards
	RAR_V5               FileType = 71  //Roshal ARchive compressed archive v5.00 onwards
	ELF                  FileType = 72  //Executable and Linkable Format
	PNG                  FileType = 73  //Image encoded in the Portable Network Graphics format
	HDF4                 FileType = 75  //Data stored in version 4 of the Hierarchical Data Format
	HDF5                 FileType = 76  //Data stored in version 5 of the Hierarchical Data Format
	COM                  FileType = 77  //CP/M 3 and higher with overlays
	JAVA_CLASS           FileType = 78  //Java class file, Mach-O Fat Binary
	UTF8_TXT             FileType = 79  //UTF-8 byte order mark, commonly seen in text files
	UTF16LE_TXT          FileType = 80  //UTF-16LE byte order mark, commonly seen in text files.
	UTF16BE_TXT          FileType = 81  //UTF-16BE byte order mark, commonly seen in text files.
	UTF32LE_TXT          FileType = 82  //UTF-32LE byte order mark for text
	UTF32BE_TXT          FileType = 83  //UTF-32BE byte order mark for text
	UTF7_TXT             FileType = 84  //UTF-7 byte order mark for text
	SCSU_TXT             FileType = 85  //SCSU byte order mark for text
	EBCDIC_TXT           FileType = 86  //UTF-EBCDIC byte order mark for text[
	MACHO_BIN32          FileType = 87  //Mach-O binary (32-bit)
	MACHO_BIN64          FileType = 88  //Mach-O binary (64-bit)
	JKS                  FileType = 89  //JKS Javakey Store
	MACHO_BIN32R         FileType = 90  //Mach-O binary (reverse byte ordering scheme, 32-bit)
	MACHO_BIN64R         FileType = 91  //Mach-O binary (reverse byte ordering scheme, 64-bit)
	PS                   FileType = 92  //PostScript document
	EPS3                 FileType = 93  //Encapsulated PostScript file version 3.0
	EPS31                FileType = 94  //Encapsulated PostScript file version 3.1
	CHM                  FileType = 95  //MS Windows HtmlHelp Data
	HLP                  FileType = 96  //Windows 3.x/95/98 Help file
	PDF                  FileType = 97  //PDF document
	ASF                  FileType = 98  //Advanced Systems Format
	MSSDI                FileType = 99  //System Deployment Image, a disk image format used by Microsoft
	OGG                  FileType = 100 //Ogg, an open source media container format
	PSD                  FileType = 101 //Photoshop Document file, Adobe Photoshop's native file format
	WAV                  FileType = 102 //Waveform Audio File Format[
	AVI                  FileType = 103 //Audio Video Interleave video format
	MP3                  FileType = 104 //MPEG-1 Layer 3 file without an ID3 tag or with an ID3v1 tag (which is appended at the end of the file)
	MP3V2                FileType = 105 //MP3 file with an ID3v2 container
	BMP                  FileType = 106 //BMP file, a bitmap format used mostly in the Windows world
	ISO                  FileType = 107 //ISO9660 CD/DVD image file
	CDI                  FileType = 108 //CD-i CD image file
	MGW                  FileType = 109 //Nintendo Game & Watch image file
	NES                  FileType = 110 //Nintendo Entertainment System image file
	D64                  FileType = 111 //Commodore 64 1541 disk image (D64 format)
	G64                  FileType = 112 //Commodore 64 1541 disk image (G64 format)
	D81                  FileType = 113 //Commodore 64 1581 disk image (D81 format)
	T64                  FileType = 114 //Commodore 64 tape image
	CRT64                FileType = 115 //Commodore 64 cartridge image
	FITS                 FileType = 116 //Flexible Image Transport System (FITS)
	FLAC                 FileType = 117 //Free Lossless Audio Codec
	MIDI                 FileType = 118 //MIDI sound file
	MSCOM                FileType = 119 //Compound File Binary Format, a container format defined by Microsoft COM. It can contain the equivalent of files and directories. It is used by Windows Installer and for documents in older versions of Microsoft Office. It can be used by other programs as well that rely on the COM and OLE API's.
	DEX                  FileType = 120 //Dalvik Executable
	VDMK                 FileType = 121 //VMDK files
)

type AnyBytesInMiddle struct {
	Prefix         []byte
	AnyBytesOffset uint8
	AnyBytesLength uint8
	Suffix         []byte
}
type AnyBytesInMiddleMask struct {
	Offset uint8
	Length uint8
}

type ByteSequence interface {
	[]byte | AnyBytesInMiddle
}
type OneOfByteSequences [][]byte

type BytesPattern interface {
	ByteSequence | OneOfByteSequences
}

type OffsetFromEof uint64
type OffsetAny bool
type OffsetMultiply []uint64
type PowerOffset struct {
	Initial uint64
	Base    uint32
	Factor  int8
}

const OFFSET_ANY OffsetAny = true

type OffsetPattern interface {
	uint64 | OffsetFromEof | OffsetAny | OffsetMultiply | PowerOffset
}

type NameExtensionPattern interface {
	string | []string | *regexp.Regexp
}

type HexSignature[S BytesPattern, O OffsetPattern, E NameExtensionPattern] struct {
	Bytes         S
	Offset        O
	NameExtension E
	Description   string
	Tag           FileType
}

var knownSignatures1 = map[FileType]HexSignature[[]byte, uint64, string]{
	SHEBANG: {
		Bytes:         []byte{0x23, 0x21},
		Offset:        0,
		NameExtension: "",
		Description:   "Script or data to be passed to the program following the shebang (#!)",
		Tag:           SHEBANG,
	},
	CLARIS_WORKS: {
		Bytes:         []byte{0x02, 0x00, 0x5a, 0x57, 0x52, 0x54, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		Offset:        0,
		NameExtension: "cwk",
		Description:   "Claris Works word processing doc",
		Tag:           CLARIS_WORKS,
	},
	LOTUS_123_V1: {
		Bytes:         []byte{0x00, 0x00, 0x02, 0x00, 0x06, 0x04, 0x06, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00},
		Offset:        0,
		NameExtension: "wk1",
		Description:   "Lotus 1-2-3 spreadsheet (v1) file",
		Tag:           LOTUS_123_V1,
	},
	LOTUS_123_V3: {
		Bytes:         []byte{0x00, 0x00, 0x1A, 0x00, 0x00, 0x10, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00},
		Offset:        0,
		NameExtension: "wk3",
		Description:   "Lotus 1-2-3 spreadsheet (v3) file",
		Tag:           LOTUS_123_V3,
	},
	LOTUS_123_V9: {
		Bytes:         []byte{0x00, 0x00, 0x1A, 0x00, 0x05, 0x10, 0x04},
		Offset:        0,
		NameExtension: "123",
		Description:   "Lotus 1-2-3 spreadsheet (v9) file",
		Tag:           LOTUS_123_V9,
	},
	AMIGA_HUNK_EXE: {
		Bytes:         []byte{0x00, 0x00, 0x03, 0xF3},
		Offset:        0,
		NameExtension: "",
		Description:   "Amiga Hunk executable file",
		Tag:           AMIGA_HUNK_EXE,
	},
	PASSWORD_GORILLA: {
		Bytes:         []byte{0x50, 0x57, 0x53, 0x33},
		Offset:        0,
		NameExtension: "psafe3",
		Description:   "Password Gorilla Password Database",
		Tag:           PASSWORD_GORILLA,
	},
	PCAPNPG: {
		Bytes:         []byte{0x0A, 0x0D, 0x0D, 0x0A},
		Offset:        0,
		NameExtension: "pcapnpg",
		Description:   "PCAP Next Generation Dump File Format",
		Tag:           PCAPNPG,
	},
	RPM: {
		Bytes:         []byte{0xED, 0xAB, 0xEE, 0xDB},
		Offset:        0,
		NameExtension: "rpm",
		Description:   "RedHat Package Manager (RPM) package",
		Tag:           RPM,
	},
	AMAZON_KINDLE_UP: {
		Bytes:         []byte{0x53, 0x50, 0x30, 0x31},
		Offset:        0,
		NameExtension: "bin",
		Description:   "Amazon Kindle Update Package",
		Tag:           AMAZON_KINDLE_UP,
	},
	DOOM_WAD: {
		Bytes:         []byte{0x49, 0x57, 0x41, 0x44},
		Offset:        0,
		NameExtension: "wad",
		Description:   "internal WAD (main resource file of Doom)",
		Tag:           DOOM_WAD,
	},
	PALM_PILOT: {
		Bytes:         []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		Offset:        11,
		NameExtension: "PDB",
		Description:   "PalmPilot Database/Document File",
		Tag:           PALM_PILOT,
	},
	PALM_DSK_CALENDAR: {
		Bytes:         []byte{0xBE, 0xBA, 0xFE, 0xCA},
		Offset:        0,
		NameExtension: "DBA",
		Description:   "Palm Desktop Calendar Archive",
		Tag:           PALM_DSK_CALENDAR,
	},
	PALM_DSK_TODO: {
		Bytes:         []byte{0x00, 0x01, 0x42, 0x44},
		Offset:        0,
		NameExtension: "DBA",
		Description:   "Palm Desktop To Do Archive",
		Tag:           PALM_DSK_TODO,
	},
	PALM_DSK_CALENDAR2: {
		Bytes:         []byte{0x00, 0x01, 0x44, 0x54},
		Offset:        0,
		NameExtension: "DBA",
		Description:   "Palm Desktop Calendar Archive",
		Tag:           PALM_DSK_CALENDAR2,
	},
	TELEGRAM_DSK: {
		Bytes:         []byte{0x54, 0x44, 0x46, 0x24},
		Offset:        0,
		NameExtension: "TDF$",
		Description:   "Telegram Desktop File",
		Tag:           TELEGRAM_DSK,
	},
	TELEGRAM_DSK_ENC: {
		Bytes:         []byte{0x54, 0x44, 0x45, 0x46},
		Offset:        0,
		NameExtension: "TDEF",
		Description:   "Telegram Desktop Encrypted File",
		Tag:           TELEGRAM_DSK_ENC,
	},
	PALM_DSK_DATA: {
		Bytes:         []byte{0x00, 0x01, 0x00, 0x00},
		Offset:        0,
		NameExtension: "",
		Description:   "Palm Desktop Data File (Access format",
		Tag:           PALM_DSK_DATA,
	},
	ICON: {
		Bytes:         []byte{0x00, 0x00, 0x01, 0x00},
		Offset:        0,
		NameExtension: "ico",
		Description:   "Computer icon encoded in ICO file format",
		Tag:           ICON,
	},
	APPLE_ICON_FORMAT: {
		Bytes:         []byte{0x69, 0x63, 0x6e, 0x73},
		Offset:        0,
		NameExtension: "icns",
		Description:   "Apple Icon Image format",
		Tag:           APPLE_ICON_FORMAT,
	},
	HEIC: {
		Bytes:         []byte{0x66, 0x74, 0x79, 0x70, 0x68, 0x65, 0x69, 0x63, 0x66, 0x74, 0x79, 0x70, 0x6d},
		Offset:        4,
		NameExtension: "heic",
		Description:   "High Efficiency Image Container (HEIC)",
		Tag:           HEIC,
	},
	LZH0: {
		Bytes:         []byte{0x2D, 0x68, 0x6C, 0x30, 0x2D},
		Offset:        2,
		NameExtension: "lzh",
		Description:   "Lempel Ziv Huffman archive file Method 0 (No compression)",
		Tag:           LZH0,
	},
	LZH5: {
		Bytes:         []byte{0x2D, 0x68, 0x6C, 0x35, 0x2D},
		Offset:        2,
		NameExtension: "lzh",
		Description:   "Lempel Ziv Huffman archive file Method 5 (8 KiB sliding window)",
		Tag:           LZH5,
	},
	AMI_BACK: {
		Bytes:         []byte{0x42, 0x41, 0x43, 0x4B, 0x4D, 0x49, 0x4B, 0x45, 0x44, 0x49, 0x53, 0x4B},
		Offset:        0,
		NameExtension: "bac",
		Description:   "AmiBack Amiga Backup data file",
		Tag:           AMI_BACK,
	},
	AMI_BACK_IDX: {
		Bytes:         []byte{0x49, 0x4E, 0x44, 0x58},
		Offset:        0,
		NameExtension: "idx",
		Description:   "AmiBack Amiga Backup index file",
		Tag:           AMI_BACK_IDX,
	},
	BPLIST: {
		Bytes:         []byte{0x62, 0x70, 0x6C, 0x69, 0x73, 0x74},
		Offset:        0,
		NameExtension: "plist",
		Description:   "Binary Property List file",
		Tag:           BPLIST,
	},
	BZ2: {
		Bytes:         []byte{0x42, 0x5A, 0x68},
		Offset:        0,
		NameExtension: "bz2",
		Description:   "Compressed file using Bzip2 algorithm",
		Tag:           BZ2,
	},
	CANON_RAW_V2: {
		Bytes:         []byte{0x49, 0x49, 0x2A, 0x00, 0x10, 0x00, 0x00, 0x00, 0x43, 0x52},
		Offset:        0,
		NameExtension: "cr2",
		Description:   "Canon RAW Format Version 2 Canon's RAW format is based on TIFF",
		Tag:           CANON_RAW_V2,
	},
	KODAK_CIN: {
		Bytes:         []byte{0x80, 0x2A, 0x5F, 0xD7},
		Offset:        0,
		NameExtension: "cin",
		Description:   "Kodak Cineon image",
		Tag:           KODAK_CIN,
	},
	RNC_V1: {
		Bytes:         []byte{0x52, 0x4E, 0x43, 0x01},
		Offset:        0,
		NameExtension: "",
		Description:   "Compressed file using Rob Northen Compression (version 1) algorithm",
		Tag:           RNC_V1,
	},
	RNC_V2: {
		Bytes:         []byte{0x52, 0x4E, 0x43, 0x02},
		Offset:        0,
		NameExtension: "",
		Description:   "Compressed file using Rob Northen Compression (version 2) algorithm",
		Tag:           RNC_V2,
	},
	NURU_IMAGE: {
		Bytes:         []byte{0x4E, 0x55, 0x52, 0x55, 0x49, 0x4D, 0x47},
		Offset:        0,
		NameExtension: "nui",
		Description:   "nuru ASCII/ANSI image files",
		Tag:           NURU_IMAGE,
	},
	NURU_PALETTE: {
		Bytes:         []byte{0x4E, 0x55, 0x52, 0x55, 0x50, 0x41, 0x4C},
		Offset:        0,
		NameExtension: "nup",
		Description:   "nuru ASCII/ANSI palette files",
		Tag:           NURU_PALETTE,
	},
	OPEN_EXR: {
		Bytes:         []byte{0x76, 0x2F, 0x31, 0x01},
		Offset:        0,
		NameExtension: "exr",
		Description:   "OpenEXR image",
		Tag:           OPEN_EXR,
	},
	BPG: {
		Bytes:         []byte{0x42, 0x50, 0x47, 0xFB},
		Offset:        0,
		NameExtension: "bpg",
		Description:   "Better Portable Graphics format",
		Tag:           BPG,
	},
	QUI: {
		Bytes:         []byte{0x71, 0x6f, 0x69, 0x66},
		Offset:        0,
		NameExtension: "qui",
		Description:   "QOI - The “Quite OK Image Format”",
		Tag:           QUI,
	},
	LZ: {
		Bytes:         []byte{0x4C, 0x5A, 0x49, 0x50},
		Offset:        0,
		NameExtension: "lz",
		Description:   "lzip compressed file",
		Tag:           LZ,
	},
	CPIO: {
		Bytes:         []byte{0x30, 0x37, 0x30, 0x37, 0x30, 0x37},
		Offset:        0,
		NameExtension: "cpio",
		Description:   "cpio archive file",
		Tag:           CPIO,
	},
	SMART_SNIFF: {
		Bytes:         []byte{0x53, 0x4D, 0x53, 0x4E, 0x46, 0x32, 0x30, 0x30},
		Offset:        0,
		NameExtension: "spp",
		Description:   "SmartSniff Packets File",
		Tag:           SMART_SNIFF,
	},
	DOS_ZM: {
		Bytes:         []byte{0x5A, 0x4D},
		Offset:        0,
		NameExtension: "exe",
		Description:   "DOS ZM executable and its descendants (rare)",
		Tag:           DOS_ZM,
	},
	RAR_V1: {
		Bytes:         []byte{0x52, 0x61, 0x72, 0x21, 0x1A, 0x07, 0x00},
		Offset:        0,
		NameExtension: "rar",
		Description:   "Roshal ARchive compressed archive v1.50 onwards",
		Tag:           RAR_V1,
	},
	RAR_V5: {
		Bytes:         []byte{0x52, 0x61, 0x72, 0x21, 0x1A, 0x07, 0x01, 0x00},
		Offset:        0,
		NameExtension: "rar",
		Description:   "Roshal ARchive compressed archive v5.00 onwards",
		Tag:           RAR_V5,
	},
	PNG: {
		Bytes:         []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A},
		Offset:        0,
		NameExtension: "png",
		Description:   "Image encoded in the Portable Network Graphics format",
		Tag:           PNG,
	},
	COM: {
		Bytes:         []byte{0xC9},
		Offset:        0,
		NameExtension: "com",
		Description:   "CP/M 3 and higher with overlays",
		Tag:           COM,
	},
	JAVA_CLASS: {
		Bytes:         []byte{0xCA, 0xFE, 0xBA, 0xBE},
		Offset:        0,
		NameExtension: "class",
		Description:   "Java class file, Mach-O Fat Binary",
		Tag:           JAVA_CLASS,
	},
	EBCDIC_TXT: {
		Bytes:         []byte{0xDD, 0x73, 0x66, 0x73},
		Offset:        0,
		NameExtension: "",
		Description:   "UTF-EBCDIC byte order mark for text",
		Tag:           EBCDIC_TXT,
	},
	JKS: {
		Bytes:         []byte{0xFE, 0xED, 0xFE, 0xED},
		Offset:        0,
		NameExtension: "",
		Description:   "JKS Javakey Store",
		Tag:           JKS,
	},
	MACHO_BIN32R: {
		Bytes:         []byte{0xCE, 0xFA, 0xED, 0xFE},
		Offset:        0,
		NameExtension: "",
		Description:   "Mach-O binary (reverse byte ordering scheme, 32-bit)",
		Tag:           MACHO_BIN32R,
	},
	MACHO_BIN64R: {
		Bytes:         []byte{0xCF, 0xFA, 0xED, 0xFE},
		Offset:        0,
		NameExtension: "",
		Description:   "Mach-O binary (reverse byte ordering scheme, 64-bit)",
		Tag:           MACHO_BIN64R,
	},
	PS: {
		Bytes:         []byte{0x25, 0x21, 0x50, 0x53},
		Offset:        0,
		NameExtension: "ps",
		Description:   "PostScript document",
		Tag:           PS,
	},
	CHM: {
		Bytes:         []byte{0x49, 0x54, 0x53, 0x46, 0x03, 0x00, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00},
		Offset:        0,
		NameExtension: "chm",
		Description:   "MS Windows HtmlHelp Data",
		Tag:           CHM,
	},
	HLP: {
		Bytes:         []byte{0x3F, 0x5F},
		Offset:        0,
		NameExtension: "hlp",
		Description:   "Windows 3.x/95/98 Help file",
		Tag:           HLP,
	},
	PDF: {
		Bytes:         []byte{0x25, 0x50, 0x44, 0x46, 0x2D},
		Offset:        0,
		NameExtension: "pdf",
		Description:   "PDF document",
		Tag:           PDF,
	},
	MSSDI: {
		Bytes:         []byte{0x24, 0x53, 0x44, 0x49, 0x30, 0x30, 0x30, 0x31},
		Offset:        0,
		NameExtension: "",
		Description:   "System Deployment Image, a disk image format used by Microsoft",
		Tag:           MSSDI,
	},
	PSD: {
		Bytes:         []byte{0x38, 0x42, 0x50, 0x53},
		Offset:        0,
		NameExtension: "psd",
		Description:   "Photoshop Document file, Adobe Photoshop's native file format",
		Tag:           PSD,
	},
	MP3V2: {
		Bytes:         []byte{0x49, 0x44, 0x33},
		Offset:        0,
		NameExtension: "mp3",
		Description:   "MP3 file with an ID3v2 container",
		Tag:           MP3V2,
	},
	CDI: {
		Bytes:         []byte{0x43, 0x44, 0x30, 0x30, 0x31},
		Offset:        0x5EAC9,
		NameExtension: "cdi",
		Description:   "CD-i CD image file",
		Tag:           CDI,
	},
	MGW: {
		Bytes:         []byte{0x6D, 0x61, 0x69, 0x6E, 0x2E, 0x62, 0x73},
		Offset:        0,
		NameExtension: "mgw",
		Description:   "Nintendo Game & Watch image file",
		Tag:           MGW,
	},
	NES: {
		Bytes:         []byte{0x4E, 0x45, 0x53},
		Offset:        0,
		NameExtension: "nes",
		Description:   "Nintendo Entertainment System image file",
		Tag:           NES,
	},
	D64: {
		Bytes:         []byte{0xA0, 0x32, 0x41, 0xA0, 0xA0, 0xA0},
		Offset:        0x165A4,
		NameExtension: "d64",
		Description:   "Commodore 64 1541 disk image (D64 format)",
		Tag:           D64,
	},
	G64: {
		Bytes:         []byte{0x47, 0x53, 0x52, 0x2D, 0x31, 0x35, 0x34, 0x31},
		Offset:        0,
		NameExtension: "g64",
		Description:   "Commodore 64 1541 disk image (G64 format)",
		Tag:           G64,
	},
	// D81: {
	// 	Bytes:         []byte{0xA0, 0x33, 0x44, 0xA0, 0xA0},
	// 	Offset:        0x61819,
	// 	NameExtension: "d81",
	// 	Description:   "Commodore 64 1581 disk image (D81 format)",
	// 	Tag:           D81,
	// },
	T64: {
		Bytes:         []byte{0x43, 0x36, 0x34, 0x20, 0x74, 0x61, 0x70, 0x65, 0x20, 0x69, 0x6D, 0x61, 0x67, 0x65, 0x20, 0x66, 0x69, 0x6C, 0x65},
		Offset:        0,
		NameExtension: "t64",
		Description:   "Commodore 64 tape image",
		Tag:           T64,
	},
	CRT64: {
		Bytes:         []byte{0x43, 0x36, 0x34, 0x20, 0x43, 0x41, 0x52, 0x54, 0x52, 0x49, 0x44, 0x47, 0x45, 0x20, 0x20, 0x20},
		Offset:        0,
		NameExtension: "crt",
		Description:   "Commodore 64 cartridge image",
		Tag:           CRT64,
	},
	FITS: {
		Bytes:         []byte{0x53, 0x49, 0x4D, 0x50, 0x4C, 0x45, 0x20, 0x20, 0x3D, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x54},
		Offset:        0,
		NameExtension: "fits",
		Description:   "Flexible Image Transport System (FITS)",
		Tag:           FITS,
	},
	FLAC: {
		Bytes:         []byte{0x66, 0x4C, 0x61, 0x43},
		Offset:        0,
		NameExtension: "flac",
		Description:   "Free Lossless Audio Codec",
		Tag:           FLAC,
	},
	DEX: {
		Bytes:         []byte{0x64, 0x65, 0x78, 0x0A, 0x30, 0x33, 0x35, 0x00},
		Offset:        0,
		NameExtension: "dex",
		Description:   "Dalvik Executable",
		Tag:           DEX,
	},
	VDMK: {
		Bytes:         []byte{0x4B, 0x44, 0x4D},
		Offset:        0,
		NameExtension: "vmdk",
		Description:   "VMDK files",
		Tag:           VDMK,
	},
}

var knownSignatures2 = map[FileType]HexSignature[[]byte, uint64, []string]{
	LOTUS_123_V4_V5: {
		Bytes:         []byte{0x00, 0x00, 0x1A, 0x00, 0x02, 0x10, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00},
		Offset:        0,
		NameExtension: []string{"wk4", "wk5"},
		Description:   "Lotus 1-2-3 spreadsheet (v4, v5) file",
		Tag:           LOTUS_123_V4_V5,
	},
	SQLITE3: {
		Bytes:         []byte{0x53, 0x51, 0x4C, 0x69, 0x74, 0x65, 0x20, 0x66, 0x6F, 0x72, 0x6D, 0x61, 0x74, 0x20, 0x33, 0x00},
		Offset:        0,
		NameExtension: []string{"sqlitedb", "sqlite", "db"},
		Description:   "SQLite Database",
		Tag:           SQLITE3,
	},
	THREE_GPP: {
		Bytes:         []byte{0x66, 0x74, 0x79, 0x70, 0x33, 0x67},
		Offset:        4,
		NameExtension: []string{"3gp", "3g2"},
		Description:   "3rd Generation Partnership Project 3GPP and 3GPP2 multimedia files",
		Tag:           THREE_GPP,
	},
	Z_LZW: {
		Bytes:         []byte{0x1F, 0x9D},
		Offset:        0,
		NameExtension: []string{"z", "tar.z"},
		Description:   "compressed file (often tar zip) using Lempel-Ziv-Welch algorithm",
		Tag:           Z_LZW,
	},
	Z_LZH: {
		Bytes:         []byte{0x1F, 0xA0},
		Offset:        0,
		NameExtension: []string{"z", "tar.z"},
		Description:   "Compressed file (often tar zip) using LZH algorithm",
		Tag:           Z_LZH,
	},
	DOS_MZ: {
		Bytes:         []byte{0x4D, 0x5A},
		Offset:        0,
		NameExtension: []string{"exe", "dll", "mui", "sys", "scr", "cpl", "ocx", "ax", "iec", "ime", "rs", "tsp", "fon", "efi"},
		Description:   "DOS MZ executable and its descendants (including NE and PE",
		Tag:           DOS_MZ,
	},
	ELF: {
		Bytes:         []byte{0x7F, 0x45, 0x4C, 0x46},
		Offset:        0,
		NameExtension: []string{"", "axf", "bin", "elf", "o", "out", "prx", "puff", "ko", "mod", "so"},
		Description:   "Executable and Linkable Format",
		Tag:           ELF,
	},
	HDF4: {
		Bytes:         []byte{0x0E, 0x03, 0x13, 0x01},
		Offset:        0,
		NameExtension: []string{"hdf4", "h4"},
		Description:   "Data stored in version 4 of the Hierarchical Data Format",
		Tag:           HDF4,
	},
	EPS3: {
		Bytes:         []byte{0x25, 0x21, 0x50, 0x53, 0x2D, 0x41, 0x64, 0x6F, 0x62, 0x65, 0x2D, 0x33, 0x2E, 0x30, 0x20, 0x45, 0x50, 0x53, 0x46, 0x2D, 0x33, 0x2E, 0x30},
		Offset:        0,
		NameExtension: []string{"eps", "epsf"},
		Description:   "Encapsulated PostScript file version 3.0",
		Tag:           EPS3,
	},
	EPS31: {
		Bytes:         []byte{0x25, 0x21, 0x50, 0x53, 0x2D, 0x41, 0x64, 0x6F, 0x62, 0x65, 0x2D, 0x33, 0x2E, 0x31, 0x20, 0x45, 0x50, 0x53, 0x46, 0x2D, 0x33, 0x2E, 0x30},
		Offset:        0,
		NameExtension: []string{"eps", "epsf"},
		Description:   "Encapsulated PostScript file version 3.1",
		Tag:           EPS31,
	},
	ASF: {
		Bytes:         []byte{0x30, 0x26, 0xB2, 0x75, 0x8E, 0x66, 0xCF, 0x11, 0xA6, 0xD9, 0x00, 0xAA, 0x00, 0x62, 0xCE, 0x6C},
		Offset:        0,
		NameExtension: []string{"asf", "wma", "wmv"},
		Description:   "Advanced Systems Format",
		Tag:           ASF,
	},
	OGG: {
		Bytes:         []byte{0x4F, 0x67, 0x67, 0x53},
		Offset:        0,
		NameExtension: []string{"ogg", "oga", "ogv"},
		Description:   "Ogg, an open source media container format",
		Tag:           OGG,
	},
	BMP: {
		Bytes:         []byte{0x42, 0x4D},
		Offset:        0,
		NameExtension: []string{"bmp", "dib"},
		Description:   "BMP file, a bitmap format used mostly in the Windows world",
		Tag:           BMP,
	},
	MIDI: {
		Bytes:         []byte{0x4D, 0x54, 0x68, 0x64},
		Offset:        0,
		NameExtension: []string{"mid", "midi"},
		Description:   "MIDI sound file",
		Tag:           MIDI,
	},
	MSCOM: {
		Bytes:         []byte{0xD0, 0xCF, 0x11, 0xE0, 0xA1, 0xB1, 0x1A, 0xE1},
		Offset:        0,
		NameExtension: []string{"doc", "xls", "ppt", "msi", "msg"},
		Description:   "Compound File Binary Format, a container format defined by Microsoft COM. It can contain the equivalent of files and directories. It is used by Windows Installer and for documents in older versions of Microsoft Office. It can be used by other programs as well that rely on the COM and OLE API's.",
		Tag:           MSCOM,
	},
}

var knownSignatures3 = map[FileType]HexSignature[OneOfByteSequences, uint64, string]{
	QUARK_EXPRESS: {
		Bytes: [][]byte{
			{0x00, 0x00, 0x49, 0x49, 0x58, 0x50, 0x52},
			{0x00, 0x00, 0x4D, 0x4D, 0x58, 0x50, 0x52},
		},
		Offset:        0,
		NameExtension: "qxd",
		Description:   "Quark Express document",
		Tag:           QUARK_EXPRESS,
	},
	LIBPCAP: {
		Bytes: [][]byte{
			{0xD4, 0xC3, 0xB2, 0xA1},
			{0xA1, 0xB2, 0xC3, 0xD4},
		},
		Offset:        0,
		NameExtension: "pcap",
		Description:   "Libpcap File Format",
		Tag:           LIBPCAP,
	},
	LIBPCAP_NS: {
		Bytes: [][]byte{
			{0x4D, 0x3C, 0xB2, 0xA1},
			{0xA1, 0xB2, 0x3C, 0x4D},
		},
		Offset:        0,
		NameExtension: "pcap",
		Description:   "Libpcap File Format (nanosecond-resolution)",
		Tag:           LIBPCAP_NS,
	},
	GIF: {
		Bytes: [][]byte{
			{0x47, 0x49, 0x46, 0x38, 0x37, 0x61},
			{0x47, 0x49, 0x46, 0x38, 0x39, 0x61},
		},
		Offset:        0,
		NameExtension: "gif",
		Description:   "Image file encoded in the Graphics Interchange Format (GIF)",
		Tag:           GIF,
	},
	SMPTE_DPX: {
		Bytes: [][]byte{
			{0x53, 0x44, 0x50, 0x58},
			{0x58, 0x50, 0x44, 0x53},
		},
		Offset:        0,
		NameExtension: "dpx",
		Description:   "SMPTE DPX image",
		Tag:           SMPTE_DPX,
	},
	UTF7_TXT: {
		Bytes: [][]byte{
			{0x2B, 0x2F, 0x76, 0x38},
			{0x2B, 0x2F, 0x76, 0x39},
			{0x2B, 0x2F, 0x76, 0x2B},
			{0x2B, 0x2F, 0x76, 0x2F},
		},
		Offset:        0,
		NameExtension: "",
		Description:   "UTF-7 byte order mark for text",
		Tag:           UTF7_TXT,
	},
	MP3: {
		Bytes: [][]byte{
			{0xFF, 0xFB},
			{0xFF, 0xF3},
			{0xFF, 0xF2},
		},
		Offset:        0,
		NameExtension: "mp3",
		Description:   "MPEG-1 Layer 3 file without an ID3 tag or with an ID3v1 tag (which is appended at the end of the file)",
		Tag:           MP3,
	},
}

var knownSignatures4 = map[FileType]HexSignature[OneOfByteSequences, uint64, []string]{
	TIFF: {
		Bytes: [][]byte{
			{0x49, 0x49, 0x2A, 0x00},
			{0x4D, 0x4D, 0x00, 0x2A},
		},
		Offset:        0,
		NameExtension: []string{"tif", "tiff"},
		Description:   "Tagged Image File Format (TIFF)",
		Tag:           TIFF,
	},
	BIG_TIFF: {
		Bytes: [][]byte{
			{0x49, 0x49, 0x2B, 0x00},
			{0x4D, 0x4D, 0x00, 0x2B},
		},
		Offset:        0,
		NameExtension: []string{"tif", "tiff"},
		Description:   "BigTIFF",
		Tag:           BIG_TIFF,
	},
	JPEG_RAW: {
		Bytes: [][]byte{
			{0xFF, 0xD8, 0xFF, 0xDB},
			{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01},
			{0xFF, 0xD8, 0xFF, 0xEE},
			{0xFF, 0xD8, 0xFF, 0xE0},
		},
		Offset:        0,
		NameExtension: []string{"jpg", "jpeg"},
		Description:   "JPEG raw or in the JFIF or Exif file format",
		Tag:           JPEG_RAW,
	},
	JPEG_2000: {
		Bytes: [][]byte{
			{0x00, 0x00, 0x00, 0x0C, 0x6A, 0x50, 0x20, 0x20, 0x0D, 0x0A, 0x87, 0x0A},
			{0xFF, 0x4F, 0xFF, 0x51},
		},
		Offset:        0,
		NameExtension: []string{"jp2", "j2k", "jpf", "jpm", "jpg2", "j2c", "jpc", "jpx", "mj2"},
		Description:   "JPEG 2000 format",
		Tag:           JPEG_2000,
	},
	ZIP: {
		Bytes: [][]byte{
			{0x50, 0x4B, 0x03, 0x04},
			{0x50, 0x4B, 0x05, 0x06},
			{0x50, 0x4B, 0x07, 0x08},
		},
		Offset:        0,
		NameExtension: []string{"zip", "aar", "apk", "docx", "epub", "ipa", "jar", "kmz", "maff", "msix", "odp", "ods", "odt", "pk3", "pk4", "pptx", "usdz", "vsdx", "xlsx", "xpi"},
		Description:   "zip file format and formats based on it, such as EPUB, JAR, ODF, OOXML",
		Tag:           ZIP,
	},
}

var knownSignatures5 = map[FileType]HexSignature[AnyBytesInMiddle, uint64, []string]{
	JPEG_RAW: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0xFF, 0xD8, 0xFF, 0xE1},
			AnyBytesOffset: 4,
			AnyBytesLength: 2,
			Suffix:         []byte{0x45, 0x78, 0x69, 0x66, 0x00, 0x00},
		},
		Offset:        0,
		NameExtension: []string{"jpg", "jpeg"},
		Description:   "JPEG raw or in the JFIF or Exif file format",
		Tag:           JPEG_RAW,
	},
	WAV: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x52, 0x49, 0x46, 0x46},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x57, 0x41, 0x56, 0x45},
		},
		Offset:        0,
		NameExtension: []string{"wav"},
		Description:   "Waveform Audio File Format",
		Tag:           WAV,
	},
	AVI: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x52, 0x49, 0x46, 0x46},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x41, 0x56, 0x49, 0x20},
		},
		Offset:        0,
		NameExtension: []string{"avi"},
		Description:   "Audio Video Interleave video format",
		Tag:           AVI,
	},
}

var knownSignatures6 = map[FileType]HexSignature[AnyBytesInMiddle, OffsetAny, []string]{
	IIF_ILBM: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x46, 0x4F, 0x52, 0x4D},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x49, 0x4C, 0x42, 0x4D},
		},
		Offset:        OFFSET_ANY,
		NameExtension: []string{"ilbm", "lbm", "ibm", "iff"},
		Description:   "IFF Interleaved Bitmap Image",
		Tag:           IIF_ILBM,
	},
	IIF_VOICE: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x46, 0x4F, 0x52, 0x4D},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x38, 0x53, 0x56, 0x58},
		},
		Offset:        OFFSET_ANY,
		NameExtension: []string{"8svx", "8sv", "svx", "snd", "iff"},
		Description:   "IFF 8-Bit Sampled Voice",
		Tag:           IIF_VOICE,
	},
	IIF_AMIGA_CB: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x46, 0x4F, 0x52, 0x4D},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x41, 0x43, 0x42, 0x4D},
		},
		Offset:        OFFSET_ANY,
		NameExtension: []string{"acbm", "iff"},
		Description:   "Amiga Contiguous Bitmap",
		Tag:           IIF_AMIGA_CB,
	},
	IIF_ANI_BMP: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x46, 0x4F, 0x52, 0x4D},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x41, 0x4E, 0x42, 0x4D},
		},
		Offset:        OFFSET_ANY,
		NameExtension: []string{"anbm", "iff"},
		Description:   "IFF Animated Bitmap",
		Tag:           IIF_ANI_BMP,
	},
	IIF_ANI_CEL: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x46, 0x4F, 0x52, 0x4D},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x41, 0x4E, 0x49, 0x4D},
		},
		Offset:        OFFSET_ANY,
		NameExtension: []string{"anim", "iff"},
		Description:   "IFF CEL Animation",
		Tag:           IIF_ANI_CEL,
	},
	IIF_FAX_IMG: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x46, 0x4F, 0x52, 0x4D},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x46, 0x41, 0x58, 0x58},
		},
		Offset:        OFFSET_ANY,
		NameExtension: []string{"faxx", "fax", "iff"},
		Description:   "IFF Facsimile Image",
		Tag:           IIF_FAX_IMG,
	},
	IIF_FT: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x46, 0x4F, 0x52, 0x4D},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x46, 0x54, 0x58, 0x54},
		},
		Offset:        OFFSET_ANY,
		NameExtension: []string{"ftxt", "iff"},
		Description:   "IFF Formatted Text",
		Tag:           IIF_FT,
	},
	IIF_MUS_SCORE_SIMPLE: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x46, 0x4F, 0x52, 0x4D},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x53, 0x4D, 0x55, 0x53},
		},
		Offset:        OFFSET_ANY,
		NameExtension: []string{"smus", "smu", "mus", "iff"},
		Description:   "IFF Simple Musical Score",
		Tag:           IIF_MUS_SCORE_SIMPLE,
	},
	IIF_MUS_SCORE: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x46, 0x4F, 0x52, 0x4D},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x43, 0x4D, 0x55, 0x53},
		},
		Offset:        OFFSET_ANY,
		NameExtension: []string{"cmus", "mus", "iff"},
		Description:   "IFF Musical Score",
		Tag:           IIF_MUS_SCORE,
	},
	IIF_YUV_IMAGE: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x46, 0x4F, 0x52, 0x4D},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x59, 0x55, 0x56, 0x4E},
		},
		Offset:        OFFSET_ANY,
		NameExtension: []string{"yuvn", "yuv", "iff"},
		Description:   "IFF YUV Image",
		Tag:           IIF_YUV_IMAGE,
	},
	IIF_AMIGA_FVM: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x46, 0x4F, 0x52, 0x4D},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x46, 0x41, 0x4E, 0x54},
		},
		Offset:        OFFSET_ANY,
		NameExtension: []string{"iff"},
		Description:   "Amiga Fantavision Movie",
		Tag:           IIF_AMIGA_FVM,
	},
	IIF_AIFF: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x46, 0x4F, 0x52, 0x4D},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x41, 0x49, 0x46, 0x46},
		},
		Offset:        OFFSET_ANY,
		NameExtension: []string{"aiff", "aif", "aifc", "snd", "iff"},
		Description:   "Audio Interchange File Format",
		Tag:           IIF_AIFF,
	},
}

// Signatures with different possible location (search offsets by power law)
var knownSignatures7 = map[FileType]HexSignature[[]byte, PowerOffset, []string]{
	HDF5: {
		Bytes: []byte{0x89, 0x48, 0x44, 0x46, 0x0D, 0x0A, 0x1A, 0x0A},
		Offset: PowerOffset{
			Initial: 0,
			Base:    512,
			Factor:  2,
		},
		NameExtension: []string{"hdf5", "h5"},
		Description:   "Data stored in version 5 of the Hierarchical Data Format",
		Tag:           HDF5,
	},
}

var knownSignatures8 = map[FileType]HexSignature[[]byte, uint64, *regexp.Regexp]{
	UTF8_TXT: {
		Bytes:         []byte{0xEF, 0xBB, 0xBF},
		Offset:        0,
		NameExtension: regexp.MustCompile("(txt|.*)"),
		Description:   "UTF-8 byte order mark, commonly seen in text files",
		Tag:           UTF8_TXT,
	},
	UTF16LE_TXT: {
		Bytes:         []byte{0xFF, 0xFE},
		Offset:        0,
		NameExtension: regexp.MustCompile("(txt|.*)"),
		Description:   "UTF-16LE byte order mark, commonly seen in text files",
		Tag:           UTF16LE_TXT,
	},
	UTF16BE_TXT: {
		Bytes:         []byte{0xFE, 0xFF},
		Offset:        0,
		NameExtension: regexp.MustCompile("(txt|.*)"),
		Description:   "UTF-16BE byte order mark, commonly seen in text files",
		Tag:           UTF16BE_TXT,
	},
	UTF32LE_TXT: {
		Bytes:         []byte{0xFF, 0xFE, 0x00, 0x00},
		Offset:        0,
		NameExtension: regexp.MustCompile("(txt|.*)"),
		Description:   "UTF-32LE byte order mark for text",
		Tag:           UTF32LE_TXT,
	},
	UTF32BE_TXT: {
		Bytes:         []byte{0x00, 0x00, 0xFE, 0xFF},
		Offset:        0,
		NameExtension: regexp.MustCompile("(txt|.*)"),
		Description:   "UTF-32BE byte order mark for text",
		Tag:           UTF32BE_TXT,
	},
	SCSU_TXT: {
		Bytes:         []byte{0x0E, 0xFE, 0xFF},
		Offset:        0,
		NameExtension: regexp.MustCompile("(txt|.*)"),
		Description:   "SCSU byte order mark for text",
		Tag:           SCSU_TXT,
	},
}

var knownSignatures9 = map[FileType]HexSignature[[]byte, OffsetMultiply, string]{
	MACHO_BIN32: {
		Bytes:         []byte{0xFE, 0xED, 0xFA, 0xCE},
		Offset:        []uint64{0, 0x1000},
		NameExtension: "",
		Description:   "Mach-O binary (32-bit)",
		Tag:           MACHO_BIN32,
	},
	MACHO_BIN64: {
		Bytes:         []byte{0xFE, 0xED, 0xFA, 0xCF},
		Offset:        []uint64{0, 0x1000},
		NameExtension: "",
		Description:   "Mach-O binary (64-bit)",
		Tag:           MACHO_BIN64,
	},
	ISO: {
		Bytes:         []byte{0x43, 0x44, 0x30, 0x30, 0x31},
		Offset:        []uint64{0x8001, 0x8801, 0x9001},
		NameExtension: "iso",
		Description:   "UTF-8 byte order mark, commonly seen in text files",
		Tag:           ISO,
	},
}

var knownSignatures10 = map[FileType]HexSignature[[]byte, uint64, []string]{
	ZERO: {
		Bytes:         []byte{0x00},
		Offset:        0,
		NameExtension: []string{"PIC", "PIF", "SEA", "YTR"},
		Description:   "IBM Storyboard bitmap file, Windows Program Information File, Mac Stuffit Self-Extracting Archive, IRIS OCR data file",
		Tag:           ZERO,
	},
}
