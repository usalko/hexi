package ft

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
	MSCOM                FileType = 119 //Compound File Binary Format, a container format defined by Microsoft COM. It can contain the equivalent of files and directories. It is used by Windows Installer and for documents in older versions of Microsoft Office. It can be used by other programs as well that rely on the COM and OLE API's
	DEX                  FileType = 120 //Dalvik Executable
	VDMK                 FileType = 121 //VMDK files
	VMWARE4VDDF          FileType = 122 //VMware 4 Virtual Disk description file (split disk)
	CRX                  FileType = 123 //Google Chrome extension or packaged app
	FH8                  FileType = 124 //FreeHand 8 document
	APPLE_WORKS5         FileType = 125 //AppleWorks 5 document
	APPLE_WORKS6         FileType = 126 //AppleWorks 6 document
	ROXIO_TOAST          FileType = 127 //Roxio Toast disc image file
	APPLE_DMG            FileType = 128 //Apple Disk Image file
	XAR                  FileType = 129 //eXtensible ARchive format
	WINDOWS_DAT          FileType = 130 //Windows Files And Settings Transfer Repository[ See also USMT 3.0 (Win XP) and USMT 4.0 (Win 7) User Guides
	NINTENDO_SYS_ROM     FileType = 131 //Nintendo Entertainment System ROM file
	TAR                  FileType = 132 //tar archive
	OAR                  FileType = 133 //OAR file archive format, where ?? is the format version
	TOX                  FileType = 134 //Open source portable voxel file
	MLV                  FileType = 135 //MLV	Magic Lantern Video file
	WINDOWS_UPDATE       FileType = 136 //Windows Update Binary Delta Compression file
	SEVEN_Z              FileType = 137 //7-Zip File Format
	GZIP                 FileType = 138 //GZIP compressed file
	XZ                   FileType = 139 //XZ compression utility using LZMA2 compression
	LZ4                  FileType = 140 //LZ4 Frame Format Remark: LZ4 block format does not offer any magic bytes
	CAB                  FileType = 141 //Microsoft Cabinet file
	MS_QUANTUM           FileType = 142 //Microsoft compressed file in Quantum format, used prior to Windows XP. File can be decompressed using Extract.exe or Expand.exe distributed with earlier versions of Windows. After compression, the last character of the original filename extension is replaced with an underscore, e.g. ‘Setup.exe’ becomes ‘Setup.ex_’
	FLIF                 FileType = 143 //Free Lossless Image Format
	MKV                  FileType = 144 //Matroska media container, including WebM
	STG                  FileType = 145 //"SEAN : Session Analysis" Training file. Also used in compatible software "Rpw : Rowperfect for Windows" and "RP3W : ROWPERFECT3 for Windows".
	DJVU                 FileType = 146 //DjVu document the following byte is either 55 (U) for single-page or 4D (M) for multi-page documents.
	DER                  FileType = 147 //DER encoded X.509 certificate
	PEM_CRT              FileType = 148 //PEM encoded X.509 certificate
	PEM_CSR              FileType = 149 //PEM encoded X.509 Certificate Signing Request
	PEM_KEY_PKCS8        FileType = 150 //PEM encoded X.509 PKCS#8 private key
	PEM_KEY_DSA          FileType = 151 //PEM encoded X.509 PKCS#1 DSA private key
	PEM_KEY_RSA          FileType = 152 //PEM encoded X.509 PKCS#1 RSA private key
	PPK2                 FileType = 153 //PuTTY private key file version 2
	PPK3                 FileType = 154 //PuTTY private key file version 3
	OPENSSH_PRIVK        FileType = 155 //OpenSSH private key file
	OPENSSH_PUBK         FileType = 156 //OpenSSH public key file
	DCM                  FileType = 157 //DICOM Medical File Format
	WOFF                 FileType = 158 //WOFF File Format 1.0
	WOFF2                FileType = 159 //WOFF File Format 2.0
	XML8                 FileType = 160 //XML (UTF-8 or other 8-bit encodings)
	XML16LE              FileType = 161 //XML (UTF-16LE)
	XML16BE              FileType = 162 //XML (UTF-16BE)
	XML32LE              FileType = 163 //XML (UTF-32LE)
	XML32BE              FileType = 164 //XML (UTF-32BE)
	XML_EBCDIC           FileType = 165 //XML (EBCDIC)
	WASM                 FileType = 166 //WebAssembly binary format
	LEP                  FileType = 167 //Lepton compressed JPEG image
	SWF                  FileType = 168 //Adobe Flash .swf
	DEB                  FileType = 169 //linux deb file
	WEBP                 FileType = 170 //Google WebP image file, where ?? ?? ?? ?? is the file size. More information on WebP File Header
	UBOOT                FileType = 171 //Das U-Boot Universal Boot Loader
	RTF                  FileType = 172 //Rich Text Format
	MSTAPE               FileType = 173 //Microsoft Tape Format
	MPEG_TS              FileType = 174 //MPEG Transport Stream (MPEG-2 Part 1)
	MPEG_PS              FileType = 175 //MPEG Program Stream (MPEG-1 Part 1 (essentially identical) and MPEG-2 Part 1)
	MPEG_VS              FileType = 176 //MPEG-1 video and MPEG-2 video (MPEG-1 Part 2 and MPEG-2 Part 2)
	MP4_ISO              FileType = 177 //ISO Base Media file (MPEG-4)
	MP4                  FileType = 178 //MPEG-4 video file
	ZLIB_NC              FileType = 179 //No Compression (no preset dictionary)
	ZLIB_BS              FileType = 180 //Best speed (no preset dictionary)
	ZLIB_DC              FileType = 181 //Default Compression (no preset dictionary)
	ZLIB_BC              FileType = 182 //Best Compression (no preset dictionary)
	ZLIB_NCPD            FileType = 183 //No Compression (with preset dictionary)
	ZLIB_BSPD            FileType = 184 //Best speed (with preset dictionary)
	ZLIB_DCPD            FileType = 185 //Default Compression (with preset dictionary)
	ZLIB_BCPD            FileType = 186 //Best Compression (with preset dictionary)
	LZFSE                FileType = 187 //Lempel-Ziv style data compression algorithm using Finite State Entropy coding. OSS by Apple.
	ORC                  FileType = 188 //Apache ORC (Optimized Row Columnar) file format
	AVRO                 FileType = 189 //Apache Avro binary file format
	RC_CFF               FileType = 190 //RCFile columnar file format
	RBXL                 FileType = 191 //Roblox place file
	PHOTOCAP_OTMPL       FileType = 192 //PhotoCap Object Templates
	PHOTOCAP_VEC         FileType = 193 //PhotoCap Vector
	PHOTOCAP_TMPL        FileType = 194 //PhotoCap Template
	PARQUET              FileType = 195 //Apache Parquet columnar file format
	EZ2                  FileType = 196 //Emulator Emaxsynth samples
	EZ3                  FileType = 197 //Emulator III synth samples
	LUA_BYTECODE         FileType = 198 //Lua bytecode[72]
	MACOS_SYMLINK        FileType = 199 //macOS file Alias (Symbolic link)
	MSZI                 FileType = 200 //Identifier	Microsoft Zone Identifier for URL Security Zones
	EML                  FileType = 201 //Email Message var5
	TDE                  FileType = 202 //Tableau Datasource
	KDB                  FileType = 203 //KDB file
	PGP                  FileType = 204 //PGP file
	ZST                  FileType = 205 //Zstandard compress
	QUICK_ZIP_RS         FileType = 206 //QuickZip rs compressed archive[79][80]
	SML                  FileType = 207 //Smile file
	PEF                  FileType = 208 //Preferred Executable Format
	SRT                  FileType = 209 //SubRip File
	VPK                  FileType = 210 //VPK file, used to store game data for some Source Engine games
	ACE                  FileType = 211 //ACE (compressed file format)[citation needed]
	ARJ                  FileType = 212 //ARJ
	IS_CAB               FileType = 213 //InstallShield CAB Archive File
	KWAJ                 FileType = 214 //Windows 3.1x Compressed File
	SZDD                 FileType = 215 //Windows 9x Compressed File
	ZOD                  FileType = 216 //Zoo (file format)
	PBM_P1               FileType = 217 //Portable bitmap ASCII
	PBM_P4               FileType = 218 //Portable bitmap binary
	PBM_P2               FileType = 219 //Portable Gray Map ASCII
	PBM_P5               FileType = 220 //Portable Gray Map binary
	PBM_P3               FileType = 221 //Pixmap ASCII
	PBM_P6               FileType = 222 //Portable Pixmap binary
	WMF                  FileType = 223 //Windows Metafile
	XCF_GIMP             FileType = 224 //XCF (file format)
	XPM                  FileType = 225 //X PixMap
	AFF                  FileType = 226 //Advanced Forensics Format
	ENCASE_EWF_V2        FileType = 227 //EnCase EWF version 2 format
	ENCASE_EWF_V1        FileType = 228 //EnCase EWF version 1 format
	QCOW                 FileType = 229 //qcow file format
	ANI                  FileType = 230 //Animated cursor
	CDA                  FileType = 231 //cda file
	QCP                  FileType = 232 //Qualcomm PureVoice file format
	SHOCKWAVE_DCR        FileType = 233 //Adobe Shockwave
	MM_DIR               FileType = 234 //Macromedia Director file format
	FLV                  FileType = 235 //Flash Video file
	VDI                  FileType = 236 //VirtualBox Virtual Hard Disk file format
	VHD                  FileType = 237 //Windows Virtual PC Virtual Hard Disk file format
	VHDX                 FileType = 238 //Windows Virtual PC Windows 8 Virtual Hard Disk file format
	ISZ                  FileType = 239 //Compressed ISO image
	DAA                  FileType = 240 //Direct Access Archive PowerISO
	EVT                  FileType = 241 //Windows Event Viewer file format
	EVT_XML              FileType = 242 //Windows Event Viewer XML file format
	SDB                  FileType = 243 //Windows customized database
	GRP                  FileType = 244 //Windows 3.x Program Manager Program Group file format
	ICM                  FileType = 245 //ICC profile
	MSREG_HIV            FileType = 246 //Windows Registry file
	MSOUTLOOK_PST        FileType = 247 //Microsoft Outlook Personal Storage Table file
	DRACO                FileType = 248 //3D model compressed with Google Draco[86]
	GRIBV1V2             FileType = 249 //Gridded data (commonly weather observations or forecasts) in the WMO GRIB or GRIB2 format
	BLENDER              FileType = 250 //Blender File Format[88]
	JXL                  FileType = 251 //Image encoded in the JPEG XL format
	TTF                  FileType = 252 //TrueType font
	OTF                  FileType = 253 //OpenType font
	MODF                 FileType = 254 //Modulefile for Environment Modules
	MSWIM                FileType = 255 //Windows Imaging Format file
	SLOB                 FileType = 256 //Slob (sorted list of blobs) is a read-only, compressed data store with dictionary-like interface
	SJD                  FileType = 257 //Serialized Java Data
	CVOCF                FileType = 258 //Creative Voice File
	AUSND                FileType = 259 //Au audio file format
	OGL_PFB              FileType = 260 //OpenGL Iris Perfomer .PFB (Performer Fast Binary)
	HAZELR               FileType = 261 //hazelrules	Noodlesoft Hazel
	FLP                  FileType = 262 //FL Studio Project File
	FLMP                 FileType = 263 //FL Studio Mobile Project File
	VORENC_DPM           FileType = 264 //Vormetric Encryption DPM Version 2.1 Header
	MSISAM               FileType = 265 //Microsoft Money file
	MSACCDB              FileType = 266 //Microsoft Access 2007 Database
	MSMDB                FileType = 267 //Microsoft Access Database
	DRW                  FileType = 268 //Micrografx vector graphic file
	DSSV2                FileType = 269 //Digital Speech Standard (Olympus, Grundig, & Phillips) v2
	DSSV3                FileType = 270 //Digital Speech Standard (Olympus, Grundig, & Phillips) v3
	ADX                  FileType = 271 //Approach index file
	ADOBE_INDD           FileType = 272 //Adobe InDesign document
	MXF                  FileType = 273 //Material Exchange Format file
	SKF                  FileType = 274 //SkinCrafter skin file
	DT2D                 FileType = 275 //DesignTools 2D Design file
	MBBTCW               FileType = 276 //MultiBit Bitcoin wallet file
	DESKMATE_DOC         FileType = 277 //DeskMate Document file
	NRI                  FileType = 278 //Nero CD Compilation
	WKS                  FileType = 279 //DeskMate Worksheet
	SIB_MUS              FileType = 280 //Sibelius Music - Score file
	DSP                  FileType = 281 //Microsoft Developer Studio project file
	AMR                  FileType = 282 //Adaptive Multi-Rate ACELP (Algebraic Code Excited Linear Prediction) Codec, commonly audio format with GSM cell phones.
	SKYPE_SILK           FileType = 283 //Audio compression format developed by Skype
	RADIANCE_HDR         FileType = 284 //Radiance High Dynamic Range image file
	VBE                  FileType = 285 //VBScript Encoded script
	CDB                  FileType = 286 //MikroTik WinBox Connection Database (Address Book)
	EXTM3U               FileType = 287 //Multimedia playlist
	M2AR                 FileType = 288 //M2 Archive, used by game developer M2
	CAPCOM_PAK           FileType = 289 //Capcom RE Engine game data archives
	CAPCOM_ARC           FileType = 290 //Capcom MT Framework game data archives
	INTERLEAF_PL         FileType = 291 //Interleaf PrinterLeaf / WorldView document format (now Broadvision QuickSilver)
	NIFTI                FileType = 292 //Single file NIfTI format, used extensively in biomedical imaging.
	NIFTI_PAIR           FileType = 293 //Header file of a .hdr/.img pair in NIfTI format, used extensively in biomedical imaging.
	RAF64                FileType = 294 //Report Builder file from Digital Metaphors
	VISRC                FileType = 295 //Resource file Visionaire 3.x Engine
	HL7                  FileType = 296 //Health Level Seven (HL7) Standard for electronic data exchange [1]
	SAP_PWRDATA_V1       FileType = 297 //SAP Power Monitor (version 1.1.0 and higher) data file
	ARC                  FileType = 298 //ARC archive file
	PGP_APUBK            FileType = 299 //Armored PGP public key
	CNT                  FileType = 300 //Windows 3.x - Windows 95 Help Contents
	VDRM                 FileType = 301 //VirtualDub
	TRID                 FileType = 302 //TrID
	SHW4                 FileType = 303 //Corel SHOW! 4.0
	SHW5                 FileType = 304 //SHOW! 5.0
	SHR5                 FileType = 305 //Corel SHOW! 5.0 player
	SHB5                 FileType = 306 //Corel SHOW! 5.0 background
	RMMP                 FileType = 307 //MacroMind Multimedia Movie or Microsoft Multimedia Movie
	ASTM_E57             FileType = 308 //ASTM E57 3D file format
	CROWD_STRIKE_CF      FileType = 309 //Crowdstrike Channel File
	UCAS                 FileType = 310 //Unreal Engine Compressed Asset Storage file
	UTOC                 FileType = 311 //Unreal Engine Table of Contents file
	COMODORE64_BIN       FileType = 312 //Commodore 64 binary file (source: *.P00 format for Power64 emulator)
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
	ByteSequence | OneOfByteSequences | []AnyBytesInMiddle
}

type OffsetFromEof uint64
type OffsetAny bool
type OffsetRange struct {
	Begin uint64
	End   uint64
}
type OffsetMultiply []uint64
type PowerOffset struct {
	Initial uint64
	Base    uint32
	Factor  int8
}
type ZeroOrAfter []FileType
type EveryNBytes struct {
	Initial uint64
	Step    uint32
}

const OFFSET_ANY OffsetAny = true

type OffsetPattern interface {
	uint64 | OffsetFromEof | OffsetAny | OffsetRange | OffsetMultiply | PowerOffset | ZeroOrAfter | EveryNBytes
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

var knownSignatures01 = map[FileType]HexSignature[[]byte, uint64, string]{
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
	D81: {
		Bytes:         []byte{0xA0, 0x33, 0x44, 0xA0, 0xA0},
		Offset:        0x61819,
		NameExtension: "d81",
		Description:   "Commodore 64 1581 disk image (D81 format)",
		Tag:           D81,
	},
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
	VMWARE4VDDF: {
		Bytes:         []byte{0x23, 0x20, 0x44, 0x69, 0x73, 0x6B, 0x20, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6F},
		Offset:        0,
		NameExtension: "vmdk",
		Description:   "VMware 4 Virtual Disk description file (split disk)",
		Tag:           VMWARE4VDDF,
	},
	CRX: {
		Bytes:         []byte{0x43, 0x72, 0x32, 0x34},
		Offset:        0,
		NameExtension: "crx",
		Description:   "Google Chrome extension or packaged app",
		Tag:           CRX,
	},
	FH8: {
		Bytes:         []byte{0x41, 0x47, 0x44, 0x33},
		Offset:        0,
		NameExtension: "fh8",
		Description:   "FreeHand 8 document",
		Tag:           FH8,
	},
	APPLE_WORKS5: {
		Bytes:         []byte{0x05, 0x07, 0x00, 0x00, 0x42, 0x4F, 0x42, 0x4F, 0x05, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
		Offset:        0,
		NameExtension: "cwk",
		Description:   "AppleWorks 5 document",
		Tag:           APPLE_WORKS5,
	},
	APPLE_WORKS6: {
		Bytes:         []byte{0x06, 0x07, 0xE1, 0x00, 0x42, 0x4F, 0x42, 0x4F, 0x06, 0x07, 0xE1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
		Offset:        0,
		NameExtension: "cwk",
		Description:   "AppleWorks 6 document",
		Tag:           APPLE_WORKS6,
	},
	XAR: {
		Bytes:         []byte{0x78, 0x61, 0x72, 0x21},
		Offset:        0,
		NameExtension: "xar",
		Description:   "eXtensible ARchive format",
		Tag:           XAR,
	},
	WINDOWS_DAT: {
		Bytes:         []byte{0x50, 0x4D, 0x4F, 0x43, 0x43, 0x4D, 0x4F, 0x43},
		Offset:        0,
		NameExtension: "dat",
		Description:   "Windows Files And Settings Transfer Repository[ See also USMT 3.0 (Win XP) and USMT 4.0 (Win 7) User Guides",
		Tag:           WINDOWS_DAT,
	},
	NINTENDO_SYS_ROM: {
		Bytes:         []byte{0x4E, 0x45, 0x53, 0x1A},
		Offset:        0,
		NameExtension: "nes",
		Description:   "Nintendo Entertainment System ROM file",
		Tag:           NINTENDO_SYS_ROM,
	},
	TOX: {
		Bytes:         []byte{0x74, 0x6F, 0x78, 0x33},
		Offset:        0,
		NameExtension: "tox",
		Description:   "Open source portable voxel file",
		Tag:           TOX,
	},
	MLV: {
		Bytes:         []byte{0x4D, 0x4C, 0x56, 0x49},
		Offset:        0,
		NameExtension: "MLV",
		Description:   "Magic Lantern Video file",
		Tag:           MLV,
	},
	SEVEN_Z: {
		Bytes:         []byte{0x37, 0x7A, 0xBC, 0xAF, 0x27, 0x1C},
		Offset:        0,
		NameExtension: "7z",
		Description:   "7-Zip File Format",
		Tag:           SEVEN_Z,
	},
	LZ4: {
		Bytes:         []byte{0x04, 0x22, 0x4D, 0x18},
		Offset:        0,
		NameExtension: "lz4",
		Description:   "LZ4 Frame Format Remark: LZ4 block format does not offer any magic bytes",
		Tag:           LZ4,
	},
	CAB: {
		Bytes:         []byte{0x4D, 0x53, 0x43, 0x46},
		Offset:        0,
		NameExtension: "cab",
		Description:   "Microsoft Cabinet file",
		Tag:           CAB,
	},
	FLIF: {
		Bytes:         []byte{0x46, 0x4C, 0x49, 0x46},
		Offset:        0,
		NameExtension: "flif",
		Description:   "Free Lossless Image Format",
		Tag:           FLIF,
	},
	STG: {
		Bytes:         []byte{0x4D, 0x49, 0x4C, 0x20},
		Offset:        0,
		NameExtension: "stg",
		Description:   "'SEAN : Session Analysis' Training file. Also used in compatible software 'Rpw : Rowperfect for Windows' and 'RP3W : ROWPERFECT3 for Windows'",
		Tag:           STG,
	},
	DER: {
		Bytes:         []byte{0x30, 0x82},
		Offset:        0,
		NameExtension: "der",
		Description:   "DER encoded X.509 certificate",
		Tag:           DER,
	},
	PPK2: {
		Bytes:         []byte{0x50, 0x75, 0x54, 0x54, 0x59, 0x2D, 0x55, 0x73, 0x65, 0x72, 0x2D, 0x4B, 0x65, 0x79, 0x2D, 0x46, 0x69, 0x6C, 0x65, 0x2D, 0x32, 0x3A},
		Offset:        0,
		NameExtension: "ppk",
		Description:   "PuTTY private key file version 2",
		Tag:           PPK2,
	},
	PPK3: {
		Bytes:         []byte{0x50, 0x75, 0x54, 0x54, 0x59, 0x2D, 0x55, 0x73, 0x65, 0x72, 0x2D, 0x4B, 0x65, 0x79, 0x2D, 0x46, 0x69, 0x6C, 0x65, 0x2D, 0x33, 0x3A},
		Offset:        0,
		NameExtension: "ppk",
		Description:   "PuTTY private key file version 3",
		Tag:           PPK3,
	},
	OPENSSH_PRIVK: {
		Bytes:         []byte{0x2D, 0x2D, 0x2D, 0x2D, 0x2D, 0x42, 0x45, 0x47, 0x49, 0x4E, 0x20, 0x4F, 0x50, 0x45, 0x4E, 0x53, 0x53, 0x48, 0x20, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x20, 0x4B, 0x45, 0x59, 0x2D, 0x2D, 0x2D, 0x2D, 0x2D},
		Offset:        0,
		NameExtension: "",
		Description:   "OpenSSH private key file",
		Tag:           OPENSSH_PRIVK,
	},
	OPENSSH_PUBK: {
		Bytes:         []byte{0x2D, 0x2D, 0x2D, 0x2D, 0x2D, 0x42, 0x45, 0x47, 0x49, 0x4E, 0x20, 0x53, 0x53, 0x48, 0x32, 0x20, 0x4B, 0x45, 0x59, 0x2D, 0x2D, 0x2D, 0x2D, 0x2D},
		Offset:        0,
		NameExtension: "pub",
		Description:   "OpenSSH public key file",
		Tag:           OPENSSH_PUBK,
	},
	DCM: {
		Bytes:         []byte{0x44, 0x49, 0x43, 0x4D},
		Offset:        128,
		NameExtension: "dcm",
		Description:   "DICOM Medical File Format",
		Tag:           DCM,
	},
	WOFF: {
		Bytes:         []byte{0x77, 0x4F, 0x46, 0x46},
		Offset:        0,
		NameExtension: "woff",
		Description:   "WOFF File Format 1.0",
		Tag:           WOFF,
	},
	WOFF2: {
		Bytes:         []byte{0x77, 0x4F, 0x46, 0x32},
		Offset:        0,
		NameExtension: "woff2",
		Description:   "WOFF File Format 2.0",
		Tag:           WOFF2,
	},
	WASM: {
		Bytes:         []byte{0x00, 0x61, 0x73, 0x6D},
		Offset:        0,
		NameExtension: "wasm",
		Description:   "WebAssembly binary format",
		Tag:           WASM,
	},
	LEP: {
		Bytes:         []byte{0xCF, 0x84, 0x01},
		Offset:        0,
		NameExtension: "lep",
		Description:   "Lepton compressed JPEG image",
		Tag:           LEP,
	},
	DEB: {
		Bytes:         []byte{0x21, 0x3C, 0x61, 0x72, 0x63, 0x68, 0x3E, 0x0A},
		Offset:        0,
		NameExtension: "deb",
		Description:   "linux deb file",
		Tag:           DEB,
	},
	UBOOT: {
		Bytes:         []byte{0x27, 0x05, 0x19, 0x56},
		Offset:        0,
		NameExtension: "",
		Description:   "U-Boot / uImage. Das U-Boot Universal Boot Loader",
		Tag:           UBOOT,
	},
	RTF: {
		Bytes:         []byte{0x7B, 0x5C, 0x72, 0x74, 0x66, 0x31},
		Offset:        0,
		NameExtension: "rtf",
		Description:   "Rich Text Format",
		Tag:           RTF,
	},
	MSTAPE: {
		Bytes:         []byte{0x54, 0x41, 0x50, 0x45},
		Offset:        0,
		NameExtension: "",
		Description:   "Microsoft Tape Format",
		Tag:           MSTAPE,
	},
	MP4_ISO: {
		Bytes:         []byte{0x66, 0x74, 0x79, 0x70, 0x69, 0x73, 0x6F, 0x6D},
		Offset:        4,
		NameExtension: "mp4",
		Description:   "ISO Base Media file (MPEG-4)",
		Tag:           MP4_ISO,
	},
	MP4: {
		Bytes:         []byte{0x66, 0x74, 0x79, 0x70, 0x4D, 0x53, 0x4E, 0x56},
		Offset:        4,
		NameExtension: "mp4",
		Description:   "MPEG-4 video file",
		Tag:           MP4,
	},
	ZLIB_NC: {
		Bytes:         []byte{0x78, 0x01},
		Offset:        0,
		NameExtension: "zlib",
		Description:   "No Compression (no preset dictionary)",
		Tag:           ZLIB_NC,
	},
	ZLIB_BS: {
		Bytes:         []byte{0x78, 0x5E},
		Offset:        0,
		NameExtension: "zlib",
		Description:   "Best speed (no preset dictionary)",
		Tag:           ZLIB_BS,
	},
	ZLIB_DC: {
		Bytes:         []byte{0x78, 0x9C},
		Offset:        0,
		NameExtension: "zlib",
		Description:   "Default Compression (no preset dictionary)",
		Tag:           ZLIB_DC,
	},
	ZLIB_BC: {
		Bytes:         []byte{0x78, 0xDA},
		Offset:        0,
		NameExtension: "zlib",
		Description:   "Best Compression (no preset dictionary)",
		Tag:           ZLIB_BC,
	},
	ZLIB_NCPD: {
		Bytes:         []byte{0x78, 0x20},
		Offset:        0,
		NameExtension: "zlib",
		Description:   "No Compression (with preset dictionary)",
		Tag:           ZLIB_NCPD,
	},
	ZLIB_BSPD: {
		Bytes:         []byte{0x78, 0x7D},
		Offset:        0,
		NameExtension: "zlib",
		Description:   "Best speed (with preset dictionary)",
		Tag:           ZLIB_BSPD,
	},
	ZLIB_DCPD: {
		Bytes:         []byte{0x78, 0xBB},
		Offset:        0,
		NameExtension: "zlib",
		Description:   "Default Compression (with preset dictionary)",
		Tag:           ZLIB_DCPD,
	},
	ZLIB_BCPD: {
		Bytes:         []byte{0x78, 0xF9},
		Offset:        0,
		NameExtension: "zlib",
		Description:   "Best Compression (with preset dictionary)",
		Tag:           ZLIB_BCPD,
	},
	LZFSE: {
		Bytes:         []byte{0x62, 0x76, 0x78, 0x32},
		Offset:        0,
		NameExtension: "lzfse",
		Description:   "LZFSE - Lempel-Ziv style data compression algorithm using Finite State Entropy coding. OSS by Apple.",
		Tag:           LZFSE,
	},
	ORC: {
		Bytes:         []byte{0x4F, 0x52, 0x43},
		Offset:        0,
		NameExtension: "orc",
		Description:   "Apache ORC (Optimized Row Columnar) file format",
		Tag:           ORC,
	},
	AVRO: {
		Bytes:         []byte{0x4F, 0x62, 0x6A, 0x01},
		Offset:        0,
		NameExtension: "avro",
		Description:   "Apache Avro binary file format",
		Tag:           AVRO,
	},
	RC_CFF: {
		Bytes:         []byte{0x53, 0x45, 0x51, 0x36},
		Offset:        0,
		NameExtension: "rc",
		Description:   "RCFile columnar file format",
		Tag:           RC_CFF,
	},
	RBXL: {
		Bytes:         []byte{0x3C, 0x72, 0x6F, 0x62, 0x6C, 0x6F, 0x78, 0x21},
		Offset:        0,
		NameExtension: "rbxl",
		Description:   "Roblox place file",
		Tag:           RBXL,
	},
	PHOTOCAP_VEC: {
		Bytes:         []byte{0x55, 0x55, 0xAA, 0xAA},
		Offset:        0,
		NameExtension: "pcv",
		Description:   "PhotoCap Vector",
		Tag:           PHOTOCAP_VEC,
	},
	PARQUET: {
		Bytes:         []byte{0x50, 0x41, 0x52, 0x31},
		Offset:        0,
		NameExtension: "",
		Description:   "Apache Parquet columnar file format",
		Tag:           PARQUET,
	},
	EZ2: {
		Bytes:         []byte{0x45, 0x4D, 0x58, 0x32},
		Offset:        0,
		NameExtension: "ez2",
		Description:   "Emulator Emaxsynth samples",
		Tag:           EZ2,
	},
	LUA_BYTECODE: {
		Bytes:         []byte{0x1B, 0x4C, 0x75, 0x61},
		Offset:        0,
		NameExtension: "luac",
		Description:   "Lua bytecode[72]",
		Tag:           LUA_BYTECODE,
	},
	EML: {
		Bytes:         []byte{0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x3A},
		Offset:        0,
		NameExtension: "eml",
		Description:   "Email Message var5",
		Tag:           EML,
	},
	TDE: {
		Bytes:         []byte{0x20, 0x02, 0x01, 0x62, 0xA0, 0x1E, 0xAB, 0x07, 0x02, 0x00, 0x00, 0x00},
		Offset:        0,
		NameExtension: "tde",
		Description:   "Tableau Datasource",
		Tag:           TDE,
	},
	KDB: {
		Bytes:         []byte{0x37, 0x48, 0x03, 0x02, 0x00, 0x00, 0x00, 0x00, 0x58, 0x35, 0x30, 0x39, 0x4B, 0x45, 0x59},
		Offset:        0,
		NameExtension: "kdb",
		Description:   "KDB file",
		Tag:           KDB,
	},
	ZST: {
		Bytes:         []byte{0x28, 0xB5, 0x2F, 0xFD},
		Offset:        0,
		NameExtension: "zst",
		Description:   "Zstandard compress",
		Tag:           ZST,
	},
	QUICK_ZIP_RS: {
		Bytes:         []byte{0x52, 0x53, 0x56, 0x4B, 0x44, 0x41, 0x54, 0x41},
		Offset:        0,
		NameExtension: "rs",
		Description:   "QuickZip rs compressed archive[79][80]",
		Tag:           QUICK_ZIP_RS,
	},
	SML: {
		Bytes:         []byte{0x3A, 0x29, 0x0A},
		Offset:        0,
		NameExtension: "sml",
		Description:   "Smile file",
		Tag:           SML,
	},
	PEF: {
		Bytes:         []byte{0x4A, 0x6F, 0x79, 0x21},
		Offset:        0,
		NameExtension: "",
		Description:   "Preferred Executable Format",
		Tag:           PEF,
	},
	SRT: {
		Bytes:         []byte{0x31, 0x0A, 0x30, 0x30},
		Offset:        0,
		NameExtension: "srt",
		Description:   "SubRip File",
		Tag:           SRT,
	},
	VPK: {
		Bytes:         []byte{0x34, 0x12, 0xAA, 0x55},
		Offset:        0,
		NameExtension: "vpk",
		Description:   "VPK file, used to store game data for some Source Engine games",
		Tag:           VPK,
	},
	ACE: {
		Bytes:         []byte{0x2A, 0x2A, 0x41, 0x43, 0x45, 0x2A, 0x2A},
		Offset:        7,
		NameExtension: "ace",
		Description:   "ACE (compressed file format)[citation needed]",
		Tag:           ACE,
	},
	ARJ: {
		Bytes:         []byte{0x60, 0xEA},
		Offset:        0,
		NameExtension: "arj",
		Description:   "ARJ",
		Tag:           ARJ,
	},
	IS_CAB: {
		Bytes:         []byte{0x49, 0x53, 0x63, 0x28},
		Offset:        0,
		NameExtension: "cab",
		Description:   "InstallShield CAB Archive File",
		Tag:           IS_CAB,
	},
	ZOD: {
		Bytes:         []byte{0x5A, 0x4F, 0x4F},
		Offset:        0,
		NameExtension: "zoo",
		Description:   "Zoo (file format)",
		Tag:           ZOD,
	},
	PBM_P1: {
		Bytes:         []byte{0x50, 0x31, 0x0A},
		Offset:        0,
		NameExtension: "pbm",
		Description:   "Portable bitmap ASCII",
		Tag:           PBM_P1,
	},
	PBM_P4: {
		Bytes:         []byte{0x50, 0x34, 0x0A},
		Offset:        0,
		NameExtension: "pbm",
		Description:   "Portable bitmap binary",
		Tag:           PBM_P4,
	},
	PBM_P2: {
		Bytes:         []byte{0x50, 0x32, 0x0A},
		Offset:        0,
		NameExtension: "pgm",
		Description:   "Portable Gray Map ASCII",
		Tag:           PBM_P2,
	},
	PBM_P5: {
		Bytes:         []byte{0x50, 0x35, 0x0A},
		Offset:        0,
		NameExtension: "pgm",
		Description:   "Portable Gray Map binary",
		Tag:           PBM_P5,
	},
	PBM_P3: {
		Bytes:         []byte{0x50, 0x33, 0x0A},
		Offset:        0,
		NameExtension: "ppm",
		Description:   "Portable Pixmap ASCII",
		Tag:           PBM_P3,
	},
	PBM_P6: {
		Bytes:         []byte{0x50, 0x36, 0x0A},
		Offset:        0,
		NameExtension: "ppm",
		Description:   "Portable Pixmap binary",
		Tag:           PBM_P6,
	},
	WMF: {
		Bytes:         []byte{0xD7, 0xCD, 0xC6, 0x9A},
		Offset:        0,
		NameExtension: "wmf",
		Description:   "Windows Metafile",
		Tag:           WMF,
	},
	XCF_GIMP: {
		Bytes:         []byte{0x67, 0x69, 0x6D, 0x70, 0x20, 0x78, 0x63, 0x66},
		Offset:        0,
		NameExtension: "xcf",
		Description:   "XCF (file format)",
		Tag:           XCF_GIMP,
	},
	XPM: {
		Bytes:         []byte{0x2F, 0x2A, 0x20, 0x58, 0x50, 0x4D, 0x20, 0x2A, 0x2F},
		Offset:        0,
		NameExtension: "xpm",
		Description:   "X PixMap",
		Tag:           XPM,
	},
	AFF: {
		Bytes:         []byte{0x41, 0x46, 0x46},
		Offset:        0,
		NameExtension: "aff",
		Description:   "Advanced Forensics Format",
		Tag:           AFF,
	},
	ENCASE_EWF_V2: {
		Bytes:         []byte{0x45, 0x56, 0x46, 0x32},
		Offset:        0,
		NameExtension: "Ex01",
		Description:   "EnCase EWF version 2 format",
		Tag:           ENCASE_EWF_V2,
	},
	ENCASE_EWF_V1: {
		Bytes:         []byte{0x45, 0x56, 0x46},
		Offset:        0,
		NameExtension: "e01",
		Description:   "EnCase EWF version 1 format",
		Tag:           ENCASE_EWF_V1,
	},
	QCOW: {
		Bytes:         []byte{0x51, 0x46, 0x49},
		Offset:        0,
		NameExtension: "qcow",
		Description:   "qcow file format",
		Tag:           QCOW,
	},
	FLV: {
		Bytes:         []byte{0x46, 0x4C, 0x56},
		Offset:        0,
		NameExtension: "flv",
		Description:   "Flash Video file",
		Tag:           FLV,
	},
	VDI: {
		Bytes:         []byte{0x3C, 0x3C, 0x3C, 0x20, 0x4F, 0x72, 0x61, 0x63, 0x6C, 0x65, 0x20, 0x56, 0x4D, 0x20, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6C, 0x42, 0x6F, 0x78, 0x20, 0x44, 0x69, 0x73, 0x6B, 0x20, 0x49, 0x6D, 0x61, 0x67, 0x65, 0x20, 0x3E, 0x3E, 0x3E},
		Offset:        0,
		NameExtension: "vdi",
		Description:   "VirtualBox Virtual Hard Disk file format",
		Tag:           VDI,
	},
	VHD: {
		Bytes:         []byte{0x63, 0x6F, 0x6E, 0x65, 0x63, 0x74, 0x69, 0x78},
		Offset:        0,
		NameExtension: "vhd",
		Description:   "Windows Virtual PC Virtual Hard Disk file format[85]",
		Tag:           VHD,
	},
	VHDX: {
		Bytes:         []byte{0x76, 0x68, 0x64, 0x78, 0x66, 0x69, 0x6C, 0x65},
		Offset:        0,
		NameExtension: "vhdx",
		Description:   "Windows Virtual PC Windows 8 Virtual Hard Disk file format",
		Tag:           VHDX,
	},
	ISZ: {
		Bytes:         []byte{0x49, 0x73, 0x5A, 0x21},
		Offset:        0,
		NameExtension: "isz",
		Description:   "Compressed ISO image",
		Tag:           ISZ,
	},
	DAA: {
		Bytes:         []byte{0x44, 0x41, 0x41},
		Offset:        0,
		NameExtension: "daa",
		Description:   "Direct Access Archive PowerISO",
		Tag:           DAA,
	},
	EVT: {
		Bytes:         []byte{0x4C, 0x66, 0x4C, 0x65},
		Offset:        0,
		NameExtension: "evt",
		Description:   "Windows Event Viewer file format",
		Tag:           EVT,
	},
	EVT_XML: {
		Bytes:         []byte{0x45, 0x6C, 0x66, 0x46, 0x69, 0x6C, 0x65},
		Offset:        0,
		NameExtension: "evtx",
		Description:   "Windows Event Viewer XML file format",
		Tag:           EVT_XML,
	},
	SDB: {
		Bytes:         []byte{0x73, 0x64, 0x62, 0x66},
		Offset:        8,
		NameExtension: "sdb",
		Description:   "Windows customized database",
		Tag:           SDB,
	},
	GRP: {
		Bytes:         []byte{0x50, 0x4D, 0x43, 0x43},
		Offset:        0,
		NameExtension: "grp",
		Description:   "Windows 3.x Program Manager Program Group file format",
		Tag:           GRP,
	},
	ICM: {
		Bytes:         []byte{0x4B, 0x43, 0x4D, 0x53},
		Offset:        0,
		NameExtension: "icm",
		Description:   "ICC profile",
		Tag:           ICM,
	},
	MSOUTLOOK_PST: {
		Bytes:         []byte{0x21, 0x42, 0x44, 0x4E},
		Offset:        0,
		NameExtension: "pst",
		Description:   "Microsoft Outlook Personal Storage Table file",
		Tag:           MSOUTLOOK_PST,
	},
	DRACO: {
		Bytes:         []byte{0x44, 0x52, 0x41, 0x43, 0x4F},
		Offset:        0,
		NameExtension: "drc",
		Description:   "3D model compressed with Google Draco",
		Tag:           DRACO,
	},
	BLENDER: {
		Bytes:         []byte{0x42, 0x4C, 0x45, 0x4E, 0x44, 0x45, 0x52},
		Offset:        0,
		NameExtension: "blend",
		Description:   "Blender File Format",
		Tag:           BLENDER,
	},
	OTF: {
		Bytes:         []byte{0x4F, 0x54, 0x54, 0x4F},
		Offset:        0,
		NameExtension: "otf",
		Description:   "OpenType font",
		Tag:           OTF,
	},
	MODF: {
		Bytes:         []byte{0x23, 0x25, 0x4D, 0x6F, 0x64, 0x75, 0x6C, 0x65},
		Offset:        0,
		NameExtension: "",
		Description:   "Modulefile for Environment Modules",
		Tag:           MODF,
	},
	SLOB: {
		Bytes:         []byte{0x21, 0x2D, 0x31, 0x53, 0x4C, 0x4F, 0x42, 0x1F},
		Offset:        0,
		NameExtension: "slob",
		Description:   "Slob (sorted list of blobs) is a read-only, compressed data store with dictionary-like interface[92]",
		Tag:           SLOB,
	},
	SJD: {
		Bytes:         []byte{0xAC, 0xED},
		Offset:        0,
		NameExtension: "",
		Description:   "Serialized Java Data",
		Tag:           SJD,
	},
	CVOCF: {
		Bytes:         []byte{0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65, 0x20, 0x56, 0x6F, 0x69, 0x63, 0x65, 0x20, 0x46, 0x69, 0x6C, 0x65, 0x1A, 0x1A, 0x00},
		Offset:        0,
		NameExtension: "voc",
		Description:   "Creative Voice file",
		Tag:           CVOCF,
	},
	OGL_PFB: {
		Bytes:         []byte{0xDB, 0x0A, 0xCE, 0x00},
		Offset:        0,
		NameExtension: "",
		Description:   "OpenGL Iris Perfomer .PFB (Performer Fast Binary)[94]",
		Tag:           OGL_PFB,
	},
	HAZELR: {
		Bytes:         []byte{0x48, 0x5a, 0x4c, 0x52, 0x00, 0x00, 0x00, 0x18},
		Offset:        0,
		NameExtension: "hazelrules",
		Description:   "Noodlesoft Hazel [95]",
		Tag:           HAZELR,
	},
	FLP: {
		Bytes:         []byte{0x46, 0x4C, 0x68, 0x64},
		Offset:        0,
		NameExtension: "flp",
		Description:   "FL Studio Project File",
		Tag:           FLP,
	},
	FLMP: {
		Bytes:         []byte{0x31, 0x30, 0x4C, 0x46},
		Offset:        0,
		NameExtension: "flm",
		Description:   "FL Studio Mobile Project File",
		Tag:           FLMP,
	},
	VORENC_DPM: {
		Bytes:         []byte{0x52, 0x4b, 0x4d, 0x43, 0x32, 0x31, 0x30},
		Offset:        0,
		NameExtension: "",
		Description:   "Vormetric Encryption DPM Version 2.1 Header[96]",
		Tag:           VORENC_DPM,
	},
	MSISAM: {
		Bytes:         []byte{0x00, 0x01, 0x00, 0x00, 0x4D, 0x53, 0x49, 0x53, 0x41, 0x4D, 0x20, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65},
		Offset:        0,
		NameExtension: "mny",
		Description:   "Microsoft Money file",
		Tag:           MSISAM,
	},
	MSACCDB: {
		Bytes:         []byte{0x00, 0x01, 0x00, 0x00, 0x53, 0x74, 0x61, 0x6E, 0x64, 0x61, 0x72, 0x64, 0x20, 0x41, 0x43, 0x45, 0x20, 0x44, 0x42},
		Offset:        0,
		NameExtension: "accdb",
		Description:   "Microsoft Access 2007 Database",
		Tag:           MSACCDB,
	},
	MSMDB: {
		Bytes:         []byte{0x00, 0x01, 0x00, 0x00, 0x53, 0x74, 0x61, 0x6E, 0x64, 0x61, 0x72, 0x64, 0x20, 0x4A, 0x65, 0x74, 0x20, 0x44, 0x42},
		Offset:        0,
		NameExtension: "mdb",
		Description:   "Microsoft Access Database",
		Tag:           MSMDB,
	},
	DRW: {
		Bytes:         []byte{0x01, 0xFF, 0x02, 0x04, 0x03, 0x02},
		Offset:        0,
		NameExtension: "drw",
		Description:   "Micrografx vector graphic file",
		Tag:           DRW,
	},
	DSSV2: {
		Bytes:         []byte{0x02, 0x64, 0x73, 0x73},
		Offset:        0,
		NameExtension: "dss",
		Description:   "Digital Speech Standard (Olympus, Grundig, & Phillips) v2",
		Tag:           DSSV2,
	},
	DSSV3: {
		Bytes:         []byte{0x03, 0x64, 0x73, 0x73},
		Offset:        0,
		NameExtension: "dss",
		Description:   "Digital Speech Standard (Olympus, Grundig, & Phillips) v3",
		Tag:           DSSV3,
	},
	ADX: {
		Bytes:         []byte{0x03, 0x00, 0x00, 0x00, 0x41, 0x50, 0x50, 0x52},
		Offset:        0,
		NameExtension: "adx",
		Description:   "Approach index file",
		Tag:           ADX,
	},
	ADOBE_INDD: {
		Bytes:         []byte{0x06, 0x06, 0xED, 0xF5, 0xD8, 0x1D, 0x46, 0xE5, 0xBD, 0x31, 0xEF, 0xE7, 0xFE, 0x74, 0xB7, 0x1D},
		Offset:        0,
		NameExtension: "indd",
		Description:   "Adobe InDesign document",
		Tag:           ADOBE_INDD,
	},
	SKF: {
		Bytes:         []byte{0x07, 0x53, 0x4B, 0x46},
		Offset:        0,
		NameExtension: "skf",
		Description:   "SkinCrafter skin file",
		Tag:           SKF,
	},
	DT2D: {
		Bytes:         []byte{0x07, 0x64, 0x74, 0x32, 0x64, 0x64, 0x74, 0x64},
		Offset:        0,
		NameExtension: "dtd",
		Description:   "DesignTools 2D Design file",
		Tag:           DT2D,
	},
	MBBTCW: {
		Bytes:         []byte{0x0A, 0x16, 0x6F, 0x72, 0x67, 0x2E, 0x62, 0x69, 0x74, 0x63, 0x6F, 0x69, 0x6E, 0x2E, 0x70, 0x72},
		Offset:        0,
		NameExtension: "wallet",
		Description:   "MultiBit Bitcoin wallet file",
		Tag:           MBBTCW,
	},
	DESKMATE_DOC: {
		Bytes:         []byte{0x0D, 0x44, 0x4F, 0x43},
		Offset:        0,
		NameExtension: "doc",
		Description:   "DeskMate Document file",
		Tag:           DESKMATE_DOC,
	},
	NRI: {
		Bytes:         []byte{0x0E, 0x4E, 0x65, 0x72, 0x6F, 0x49, 0x53, 0x4F},
		Offset:        0,
		NameExtension: "nri",
		Description:   "Nero CD Compilation",
		Tag:           NRI,
	},
	WKS: {
		Bytes:         []byte{0x0E, 0x57, 0x4B, 0x53},
		Offset:        0,
		NameExtension: "wks",
		Description:   "DeskMate Worksheet",
		Tag:           WKS,
	},
	SIB_MUS: {
		Bytes:         []byte{0x0F, 0x53, 0x49, 0x42, 0x45, 0x4C, 0x49, 0x55, 0x53},
		Offset:        0,
		NameExtension: "sib",
		Description:   "Sibelius Music - Score file",
		Tag:           SIB_MUS,
	},
	DSP: {
		Bytes:         []byte{0x23, 0x20, 0x4D, 0x69, 0x63, 0x72, 0x6F, 0x73, 0x6F, 0x66, 0x74, 0x20, 0x44, 0x65, 0x76, 0x65, 0x6C, 0x6F, 0x70, 0x65, 0x72, 0x20, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6F},
		Offset:        0,
		NameExtension: "dsp",
		Description:   "Microsoft Developer Studio project file",
		Tag:           DSP,
	},
	AMR: {
		Bytes:         []byte{0x23, 0x21, 0x41, 0x4D, 0x52},
		Offset:        0,
		NameExtension: "amr",
		Description:   "Adaptive Multi-Rate ACELP (Algebraic Code Excited Linear Prediction) Codec, commonly audio format with GSM cell phones.",
		Tag:           AMR,
	},
	SKYPE_SILK: {
		Bytes:         []byte{0x23, 0x21, 0x53, 0x49, 0x4C, 0x4B, 0x0A},
		Offset:        0,
		NameExtension: "sil",
		Description:   "Audio compression format developed by Skype",
		Tag:           SKYPE_SILK,
	},
	RADIANCE_HDR: {
		Bytes:         []byte{0x23, 0x3F, 0x52, 0x41, 0x44, 0x49, 0x41, 0x4E, 0x43, 0x45, 0x0A},
		Offset:        0,
		NameExtension: "hdr",
		Description:   "Radiance High Dynamic Range image file",
		Tag:           RADIANCE_HDR,
	},
	VBE: {
		Bytes:         []byte{0x23, 0x40, 0x7E, 0x5E},
		Offset:        0,
		NameExtension: "vbe",
		Description:   "VBScript Encoded script",
		Tag:           VBE,
	},
	CDB: {
		Bytes:         []byte{0x0D, 0xF0, 0x1D, 0xC0},
		Offset:        0,
		NameExtension: "cdb",
		Description:   "MikroTik WinBox Connection Database (Address Book)",
		Tag:           CDB,
	},
	M2AR: {
		Bytes:         []byte{0x6D, 0x64, 0x66, 0x00},
		Offset:        0,
		NameExtension: "m",
		Description:   "M2 Archive, used by game developer M2",
		Tag:           M2AR,
	},
	CAPCOM_PAK: {
		Bytes:         []byte{0x4B, 0x50, 0x4B, 0x41},
		Offset:        0,
		NameExtension: "pak",
		Description:   "Capcom RE Engine game data archives",
		Tag:           CAPCOM_PAK,
	},
	CAPCOM_ARC: {
		Bytes:         []byte{0x41, 0x52, 0x43},
		Offset:        0,
		NameExtension: "arc",
		Description:   "Capcom MT Framework game data archives",
		Tag:           CAPCOM_ARC,
	},
	INTERLEAF_PL: {
		Bytes:         []byte{0xD0, 0x4F, 0x50, 0x53},
		Offset:        0,
		NameExtension: "pl",
		Description:   "Interleaf PrinterLeaf / WorldView document format (now Broadvision QuickSilver)",
		Tag:           INTERLEAF_PL,
	},
	NIFTI: {
		Bytes:         []byte{0x6E, 0x2B, 0x31, 0x00},
		Offset:        344,
		NameExtension: "nii",
		Description:   "Single file NIfTI format, used extensively in biomedical imaging.",
		Tag:           NIFTI,
	},
	NIFTI_PAIR: {
		Bytes:         []byte{0x6E, 0x69, 0x31, 0x00},
		Offset:        344,
		NameExtension: "hdr",
		Description:   "Header file of a .hdr/.img pair in NIfTI format, used extensively in biomedical imaging.",
		Tag:           NIFTI_PAIR,
	},
	RAF64: {
		Bytes:         []byte{0x52, 0x41, 0x46, 0x36, 0x34},
		Offset:        0,
		NameExtension: "",
		Description:   "Report Builder file from Digital Metaphors",
		Tag:           RAF64,
	},
	VISRC: {
		Bytes:         []byte{0x56, 0x49, 0x53, 0x33},
		Offset:        0,
		NameExtension: "",
		Description:   "Resource file Visionaire 3.x Engine",
		Tag:           VISRC,
	},
	HL7: {
		Bytes:         []byte{0x4D, 0x53, 0x48, 0x7C, 0x42, 0x53, 0x48, 0x7C},
		Offset:        0,
		NameExtension: "hl7",
		Description:   "Health Level Seven (HL7) Standard for electronic data exchange [1]",
		Tag:           HL7,
	},
	SAP_PWRDATA_V1: {
		Bytes:         []byte{0x70, 0x77, 0x72, 0x64, 0x61, 0x74, 0x61},
		Offset:        0,
		NameExtension: "pwrdata",
		Description:   "SAP Power Monitor (version 1.1.0 and higher) data file",
		Tag:           SAP_PWRDATA_V1,
	},
	ARC: {
		Bytes:         []byte{0x1a, 0x08},
		Offset:        0,
		NameExtension: "arc",
		Description:   "ARC archive file",
		Tag:           ARC,
	},
	PGP_APUBK: {
		Bytes:         []byte{0x2d, 0x2d, 0x2d, 0x2d, 0x2d, 0x42, 0x45, 0x47, 0x49, 0x4e, 0x20, 0x50, 0x47, 0x50, 0x20, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x20, 0x4b, 0x45, 0x49, 0x20, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x2d, 0x2d, 0x2d, 0x2d, 0x2d},
		Offset:        0,
		NameExtension: "asc",
		Description:   "Armored PGP public key",
		Tag:           PGP_APUBK,
	},
	CNT: {
		Bytes:         []byte{0x3a, 0x42, 0x61, 0x73, 0x65, 0x20},
		Offset:        0,
		NameExtension: "cnt",
		Description:   "Windows 3.x - Windows 95 Help Contents",
		Tag:           CNT,
	},
	ASTM_E57: {
		Bytes:         []byte{0x41, 0x53, 0x54, 0x4d, 0x2d, 0x45, 0x35, 0x37},
		Offset:        0,
		NameExtension: "e57",
		Description:   "ASTM E57 3D file format",
		Tag:           ASTM_E57,
	},
	CROWD_STRIKE_CF: {
		Bytes:         []byte{0xaa, 0xaa, 0xaa, 0xaa},
		Offset:        0,
		NameExtension: "sys",
		Description:   "Crowdstrike Channel File",
		Tag:           CROWD_STRIKE_CF,
	},
	UCAS: {
		Bytes:         []byte{0x8C, 0x0A, 0x00},
		Offset:        0,
		NameExtension: "ucas",
		Description:   "Unreal Engine Compressed Asset Storage file",
		Tag:           UCAS,
	},
	UTOC: {
		Bytes:         []byte{0x2D, 0x3D, 0x3D, 0x2D, 0x2D, 0x3D, 0x3D, 0x2D, 0x2D, 0x3D, 0x3D, 0x2D, 0x2D, 0x3D, 0x3D, 0x2D},
		Offset:        0,
		NameExtension: "utoc",
		Description:   "Unreal Engine Table of Contents file",
		Tag:           UTOC,
	},
}

var knownSignatures02 = map[FileType]HexSignature[[]byte, uint64, []string]{
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
	GZIP: {
		Bytes:         []byte{0x1F, 0x8B},
		Offset:        0,
		NameExtension: []string{"gz", "tar.gz"},
		Description:   "GZIP compressed file",
		Tag:           GZIP,
	},
	XZ: {
		Bytes:         []byte{0xFD, 0x37, 0x7A, 0x58, 0x5A, 0x00},
		Offset:        0,
		NameExtension: []string{"xz", "tar.xz"},
		Description:   "XZ compression utility using LZMA2 compression",
		Tag:           XZ,
	},
	MKV: {
		Bytes:         []byte{0x1A, 0x45, 0xDF, 0xA3},
		Offset:        0,
		NameExtension: []string{"mkv", "mka", "mks", "mk3d", "webm"},
		Description:   "Matroska media container, including WebM",
		Tag:           MKV,
	},
	PEM_CRT: {
		Bytes:         []byte{0x2D, 0x2D, 0x2D, 0x2D, 0x2D, 0x42, 0x45, 0x47, 0x49, 0x4E, 0x20, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x2D, 0x2D, 0x2D, 0x2D, 0x2D},
		Offset:        0,
		NameExtension: []string{"crt", "pem"},
		Description:   "PEM encoded X.509 certificate",
		Tag:           PEM_CRT,
	},
	PEM_CSR: {
		Bytes:         []byte{0x2D, 0x2D, 0x2D, 0x2D, 0x2D, 0x42, 0x45, 0x47, 0x49, 0x4E, 0x20, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x20, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x2D, 0x2D, 0x2D, 0x2D, 0x2D},
		Offset:        0,
		NameExtension: []string{"csr", "pem"},
		Description:   "PEM encoded X.509 Certificate Signing Request",
		Tag:           PEM_CSR,
	},
	PEM_KEY_PKCS8: {
		Bytes:         []byte{0x2D, 0x2D, 0x2D, 0x2D, 0x2D, 0x42, 0x45, 0x47, 0x49, 0x4E, 0x20, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x20, 0x4B, 0x45, 0x59, 0x2D, 0x2D, 0x2D, 0x2D, 0x2D},
		Offset:        0,
		NameExtension: []string{"key", "pem"},
		Description:   "PEM encoded X.509 PKCS#8 private key",
		Tag:           PEM_KEY_PKCS8,
	},
	PEM_KEY_DSA: {
		Bytes:         []byte{0x2D, 0x2D, 0x2D, 0x2D, 0x2D, 0x42, 0x45, 0x47, 0x49, 0x4E, 0x20, 0x44, 0x53, 0x41, 0x20, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x20, 0x4B, 0x45, 0x59, 0x2D, 0x2D, 0x2D, 0x2D, 0x2D},
		Offset:        0,
		NameExtension: []string{"key", "pem"},
		Description:   "PEM encoded X.509 PKCS#1 DSA private key",
		Tag:           PEM_KEY_DSA,
	},
	PEM_KEY_RSA: {
		Bytes:         []byte{0x2D, 0x2D, 0x2D, 0x2D, 0x2D, 0x42, 0x45, 0x47, 0x49, 0x4E, 0x20, 0x52, 0x45, 0x41, 0x20, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x20, 0x4B, 0x45, 0x59, 0x2D, 0x2D, 0x2D, 0x2D, 0x2D},
		Offset:        0,
		NameExtension: []string{"key", "pem"},
		Description:   "PEM encoded X.509 PKCS#1 RSA private key",
		Tag:           PEM_KEY_RSA,
	},
	MPEG_PS: {
		Bytes:         []byte{0x00, 0x00, 0x01, 0xBA},
		Offset:        0,
		NameExtension: []string{"m2p", "vob", "mpg", "mpeg"},
		Description:   "MPEG Program Stream (MPEG-1 Part 1 (essentially identical) and MPEG-2 Part 1)",
		Tag:           MPEG_PS,
	},
	MPEG_VS: {
		Bytes:         []byte{0x00, 0x00, 0x01, 0xB3},
		Offset:        0,
		NameExtension: []string{"mpg", "mpeg"},
		Description:   "MPEG-1 video and MPEG-2 video (MPEG-1 Part 2 and MPEG-2 Part 2)",
		Tag:           MPEG_VS,
	},
	PHOTOCAP_OTMPL: {
		Bytes:         []byte{0x65, 0x87, 0x78, 0x56},
		Offset:        0,
		NameExtension: []string{"p25", "obt"},
		Description:   "PhotoCap Object Templates",
		Tag:           PHOTOCAP_OTMPL,
	},
	PHOTOCAP_TMPL: {
		Bytes:         []byte{0x78, 0x56, 0x34},
		Offset:        0,
		NameExtension: []string{"pbt", "pdt", "pea", "peb", "pet", "pgt", "pict", "pjt", "pkt", "pmt"},
		Description:   "PhotoCap Template",
		Tag:           PHOTOCAP_TMPL,
	},
	EZ3: {
		Bytes:         []byte{0x45, 0x4D, 0x55, 0x33},
		Offset:        0,
		NameExtension: []string{"ez3", "iso"},
		Description:   "Emulator III synth samples",
		Tag:           EZ3,
	},
	MSREG_HIV: {
		Bytes:         []byte{0x72, 0x65, 0x67, 0x66},
		Offset:        0,
		NameExtension: []string{"dat", "hiv"},
		Description:   "Windows Registry file",
		Tag:           MSREG_HIV,
	},
	GRIBV1V2: {
		Bytes:         []byte{0x47, 0x52, 0x49, 0x42},
		Offset:        0,
		NameExtension: []string{"grib", "grib2"},
		Description:   "Gridded data (commonly weather observations or forecasts) in the WMO GRIB or GRIB2 format[87]",
		Tag:           GRIBV1V2,
	},
	TTF: {
		Bytes:         []byte{0x00, 0x01, 0x00, 0x00, 0x00},
		Offset:        0,
		NameExtension: []string{"ttf", "tte", "dfont"},
		Description:   "TrueType font",
		Tag:           TTF,
	},
	MSWIM: {
		Bytes:         []byte{0x4D, 0x53, 0x57, 0x49, 0x4D, 0x00, 0x00, 0x00, 0xD0, 0x00, 0x00, 0x00, 0x00},
		Offset:        0,
		NameExtension: []string{"wim", "swm", "esd"},
		Description:   "Windows Imaging Format file",
		Tag:           MSWIM,
	},
	AUSND: {
		Bytes:         []byte{0x2E, 0x73, 0x6E, 0x64},
		Offset:        0,
		NameExtension: []string{"au", "snd"},
		Description:   "Au audio file format",
		Tag:           AUSND,
	},
	EXTM3U: {
		Bytes:         []byte{0x23, 0x45, 0x58, 0x54, 0x4D, 0x33, 0x55},
		Offset:        0,
		NameExtension: []string{"m3u", "m3u8"},
		Description:   "Multimedia playlist",
		Tag:           EXTM3U,
	},
}

var knownSignatures03 = map[FileType]HexSignature[OneOfByteSequences, uint64, string]{
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
	ROXIO_TOAST: {
		Bytes: [][]byte{
			{0x45, 0x52, 0x02, 0x00, 0x00, 0x00},
			{0x8B, 0x45, 0x52, 0x02, 0x00, 0x00, 0x00},
		},
		Offset:        0,
		NameExtension: "toast",
		Description:   "Roxio Toast disc image file",
		Tag:           ROXIO_TOAST,
	},
	TAR: {
		Bytes: [][]byte{
			{0x75, 0x73, 0x74, 0x61, 0x72, 0x00, 0x30},
			{0x30, 0x75, 0x73, 0x74, 0x61, 0x72, 0x20, 0x20, 0x00},
		},
		Offset:        257,
		NameExtension: "tar",
		Description:   "tar archive",
		Tag:           TAR,
	},
	WINDOWS_UPDATE: {
		Bytes: [][]byte{
			{0x44, 0x43, 0x4D, 0x01, 0x50, 0x41, 0x33, 0x30},
			{0x50, 0x41, 0x33, 0x30},
		},
		Offset:        0,
		NameExtension: "",
		Description:   "Windows Update Binary Delta Compression file",
		Tag:           WINDOWS_UPDATE,
	},
	SWF: {
		Bytes: [][]byte{
			{0x43, 0x57, 0x53},
			{0x46, 0x57, 0x53},
		},
		Offset:        0,
		NameExtension: "swf",
		Description:   "Adobe Flash .swf",
		Tag:           SWF,
	},
	MACOS_SYMLINK: {
		Bytes: [][]byte{
			{0x62, 0x6F, 0x6F, 0x6B, 0x00, 0x00, 0x00, 0x00},
			{0x6D, 0x61, 0x72, 0x6B, 0x00, 0x00, 0x00, 0x00},
		},
		Offset:        0,
		NameExtension: "alias",
		Description:   "macOS file Alias (Symbolic link)",
		Tag:           MACOS_SYMLINK,
	},
	MSZI: {
		Bytes: [][]byte{
			{0x5B, 0x5A, 0x6F, 0x6E, 0x65, 0x54, 0x72, 0x61},
			{0x6E, 0x73, 0x66, 0x65, 0x72, 0x5D},
		},
		Offset:        0,
		NameExtension: "Identifier",
		Description:   "Microsoft Zone Identifier for URL Security Zones",
		Tag:           MSZI,
	},
	JXL: {
		Bytes: [][]byte{
			{0x00, 0x00, 0x00, 0x0C, 0x4A, 0x58, 0x4C, 0x20, 0x0D, 0x0A, 0x87, 0x0A},
			{0xFF, 0x0A},
		},
		Offset:        0,
		NameExtension: "jxl",
		Description:   "Image encoded in the JPEG XL format",
		Tag:           JXL,
	},
}

var knownSignatures04 = map[FileType]HexSignature[OneOfByteSequences, uint64, []string]{
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

var knownSignatures05 = map[FileType]HexSignature[AnyBytesInMiddle, uint64, []string]{
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
	DJVU: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x41, 0x54, 0x26, 0x54, 0x46, 0x4F, 0x52, 0x4D},
			AnyBytesOffset: 8,
			AnyBytesLength: 4,
			Suffix:         []byte{0x44, 0x4A, 0x56},
		},
		Offset:        0,
		NameExtension: []string{"djvu", "djv"},
		Description:   "DjVu document the following byte is either 55 (U) for single-page or 4D (M) for multi-page documents.",
		Tag:           DJVU,
	},
	WEBP: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x52, 0x49, 0x46, 0x46},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x57, 0x45, 0x42, 0x50},
		},
		Offset:        0,
		NameExtension: []string{"webp"},
		Description:   "Google WebP image file, where ?? ?? ?? ?? is the file size. More information on WebP File Header",
		Tag:           WEBP,
	},
	PGP: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x85},
			AnyBytesOffset: 1,
			AnyBytesLength: 2,
			Suffix:         []byte{0x03},
		},
		Offset:        0,
		NameExtension: []string{"pgp"},
		Description:   "PGP file",
		Tag:           PGP,
	},
	VDRM: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x52, 0x49, 0x46, 0x46},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x56, 0x44, 0x52, 0x4d},
		},
		Offset:        0,
		NameExtension: []string{"vdr"},
		Description:   "VirtualDub",
		Tag:           VDRM,
	},
	TRID: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x52, 0x59, 0x46, 0x46},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x54, 0x52, 0x49, 0x44},
		},
		Offset:        0,
		NameExtension: []string{"trd"},
		Description:   "TrID",
		Tag:           TRID,
	},
	SHW4: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x52, 0x49, 0x46, 0x46},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x73, 0x68, 0x77, 0x34},
		},
		Offset:        0,
		NameExtension: []string{"shw"},
		Description:   "Corel SHOW! 4.0",
		Tag:           SHW4,
	},
	SHW5: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x52, 0x49, 0x46, 0x46},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x73, 0x68, 0x77, 0x35},
		},
		Offset:        0,
		NameExtension: []string{"shw"},
		Description:   "Corel SHOW! 5.0",
		Tag:           SHW5,
	},
	SHR5: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x52, 0x49, 0x46, 0x46},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x73, 0x68, 0x72, 0x35},
		},
		Offset:        0,
		NameExtension: []string{"shr"},
		Description:   "Corel SHOW! 5.0 player",
		Tag:           SHR5,
	},
	SHB5: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x52, 0x49, 0x46, 0x46},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x73, 0x68, 0x62, 0x35},
		},
		Offset:        0,
		NameExtension: []string{"shb"},
		Description:   "Corel SHOW! 5.0 background",
		Tag:           SHB5,
	},
	RMMP: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x52, 0x49, 0x46, 0x46},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x52, 0x4d, 0x4d, 0x50},
		},
		Offset:        0,
		NameExtension: []string{"mmm"},
		Description:   "MacroMind Multimedia Movie or Microsoft Multimedia Movie",
		Tag:           RMMP,
	},
	OAR: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x4F, 0x41, 0x52},
			AnyBytesOffset: 3,
			AnyBytesLength: 1,
			Suffix:         []byte{},
		},
		Offset:        0,
		NameExtension: []string{"oar"},
		Description:   "OAR file archive format, where ?? is the format version",
		Tag:           OAR,
	},
	ANI: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x52, 0x49, 0x46, 0x46},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x41, 0x43, 0x4F, 0x4E},
		},
		Offset:        0,
		NameExtension: []string{"ani"},
		Description:   "Animated cursor",
		Tag:           ANI,
	},
	CDA: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x52, 0x49, 0x46, 0x46},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x43, 0x44, 0x44, 0x41},
		},
		Offset:        0,
		NameExtension: []string{"cda"},
		Description:   ".cda file",
		Tag:           CDA,
	},
	QCP: {
		Bytes: AnyBytesInMiddle{
			Prefix:         []byte{0x52, 0x49, 0x46, 0x46},
			AnyBytesOffset: 4,
			AnyBytesLength: 4,
			Suffix:         []byte{0x51, 0x4C, 0x43, 0x4D},
		},
		Offset:        0,
		NameExtension: []string{"qcp"},
		Description:   "Qualcomm PureVoice file format",
		Tag:           QCP,
	},
}

var knownSignatures06 = map[FileType]HexSignature[AnyBytesInMiddle, OffsetAny, []string]{
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

var knownSignatures07 = map[FileType]HexSignature[[]AnyBytesInMiddle, uint64, []string]{
	SHOCKWAVE_DCR: {
		Bytes: []AnyBytesInMiddle{
			{
				Prefix:         []byte{0x52, 0x49, 0x46, 0x58},
				AnyBytesOffset: 4,
				AnyBytesLength: 4,
				Suffix:         []byte{0x46, 0x47, 0x44, 0x4D},
			},
			{
				Prefix:         []byte{0x58, 0x46, 0x49, 0x52},
				AnyBytesOffset: 4,
				AnyBytesLength: 4,
				Suffix:         []byte{0x4D, 0x44, 0x47, 0x46},
			},
		},
		Offset:        0,
		NameExtension: []string{"dcr"},
		Description:   "Adobe Shockwave",
		Tag:           SHOCKWAVE_DCR,
	},
	MM_DIR: {
		Bytes: []AnyBytesInMiddle{
			{
				Prefix:         []byte{0x52, 0x49, 0x46, 0x58},
				AnyBytesOffset: 4,
				AnyBytesLength: 4,
				Suffix:         []byte{0x4D, 0x56, 0x39, 0x33},
			},
			{
				Prefix:         []byte{0x58, 0x46, 0x49, 0x52},
				AnyBytesOffset: 4,
				AnyBytesLength: 4,
				Suffix:         []byte{0x33, 0x39, 0x56, 0x4D},
			},
		},
		Offset:        0,
		NameExtension: []string{"dir", "dxr", "drx"},
		Description:   "Macromedia Director file format",
		Tag:           MM_DIR,
	},
}

// Signatures with different possible location (search offsets by power law)
var knownSignatures08 = map[FileType]HexSignature[[]byte, PowerOffset, []string]{
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

var knownSignatures09 = map[FileType]HexSignature[[]byte, uint64, *regexp.Regexp]{
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
	KWAJ: {
		Bytes:         []byte{0x4B, 0x57, 0x41, 0x4A},
		Offset:        0,
		NameExtension: regexp.MustCompile(".._"),
		Description:   "Windows 3.1x Compressed File",
		Tag:           KWAJ,
	},
	SZDD: {
		Bytes:         []byte{0x53, 0x5A, 0x44, 0x44},
		Offset:        0,
		NameExtension: regexp.MustCompile(".._"),
		Description:   "Windows 9x Compressed File",
		Tag:           SZDD,
	},
	MS_QUANTUM: {
		Bytes:         []byte{0x53, 0x5A, 0x44, 0x44, 0x88, 0xF0, 0x27, 0x33},
		Offset:        0,
		NameExtension: regexp.MustCompile(".._"),
		Description:   "Microsoft compressed file in Quantum format, used prior to Windows XP. File can be decompressed using Extract.exe or Expand.exe distributed with earlier versions of Windows. After compression, the last character of the original filename extension is replaced with an underscore, e.g. ‘Setup.exe’ becomes ‘Setup.ex_’",
		Tag:           MS_QUANTUM,
	},
	COMODORE64_BIN: {
		Bytes:         []byte{0x43, 0x36, 0x34, 0x46, 0x69, 0x6C, 0x65, 0x00},
		Offset:        0,
		NameExtension: regexp.MustCompile(`P\d\d`),
		Description:   "Commodore 64 binary file (source: *.P00 format for Power64 emulator)",
		Tag:           COMODORE64_BIN,
	},
}

var knownSignatures10 = map[FileType]HexSignature[[]byte, OffsetMultiply, string]{
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

var knownSignatures11 = map[FileType]HexSignature[[]byte, ZeroOrAfter, string]{
	XML8: {
		Bytes:         []byte{0x3C, 0x3F, 0x78, 0x6D, 0x6C, 0x20},
		Offset:        []FileType{UTF8_TXT},
		NameExtension: "xml",
		Description:   "XML (UTF-8 or other 8-bit encodings)",
		Tag:           XML8,
	},
	XML16LE: {
		Bytes:         []byte{0x3C, 0x00, 0x3F, 0x00, 0x78, 0x00, 0x6D, 0x00, 0x6C, 0x00, 0x20},
		Offset:        []FileType{UTF8_TXT},
		NameExtension: "xml",
		Description:   "XML (UTF-16LE)",
		Tag:           XML16LE,
	},
	XML16BE: {
		Bytes:         []byte{0x00, 0x3C, 0x00, 0x3F, 0x00, 0x78, 0x00, 0x6D, 0x00, 0x6C, 0x00, 0x20},
		Offset:        []FileType{UTF8_TXT},
		NameExtension: "xml",
		Description:   "XML (UTF-16BE)",
		Tag:           XML16BE,
	},
	XML32LE: {
		Bytes:         []byte{0x3C, 0x00, 0x00, 0x00, 0x3F, 0x00, 0x00, 0x00, 0x78, 0x00, 0x00, 0x00, 0x6D, 0x00, 0x00, 0x00, 0x6C, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00},
		Offset:        []FileType{UTF8_TXT},
		NameExtension: "xml",
		Description:   "XML (UTF-32LE)",
		Tag:           XML32LE,
	},
	XML32BE: {
		Bytes:         []byte{0x00, 0x00, 0x00, 0x3C, 0x00, 0x00, 0x00, 0x3F, 0x00, 0x00, 0x00, 0x78, 0x00, 0x00, 0x00, 0x6D, 0x00, 0x00, 0x00, 0x6C, 0x00, 0x00, 0x00, 0x20},
		Offset:        []FileType{UTF8_TXT},
		NameExtension: "xml",
		Description:   "XML (UTF-32BE)",
		Tag:           XML32BE,
	},
	XML_EBCDIC: {
		Bytes:         []byte{0x4C, 0x6F, 0xA7, 0x94, 0x93, 0x40},
		Offset:        []FileType{UTF8_TXT},
		NameExtension: "xml",
		Description:   "XML (EBCDIC)",
		Tag:           XML_EBCDIC,
	},
}

var knownSignatures12 = map[FileType]HexSignature[[]byte, EveryNBytes, []string]{
	MPEG_TS: {
		Bytes: []byte{0x47},
		Offset: EveryNBytes{
			Initial: 0,
			Step:    188, // 0, 0xBC, 0x178
		},
		NameExtension: []string{"ts", "tsv", "tsa", "mpg", "mpeg"},
		Description:   "MPEG Transport Stream (MPEG-2 Part 1)",
		Tag:           MPEG_TS,
	},
}

var knownSignatures13 = map[FileType]HexSignature[[]byte, OffsetFromEof, string]{
	APPLE_DMG: {
		Bytes:         []byte{0x6B, 0x6F, 0x6C, 0x79},
		Offset:        512,
		NameExtension: "dmg",
		Description:   "Apple Disk Image file",
		Tag:           APPLE_DMG,
	},
}

var knownSignatures14 = map[FileType]HexSignature[[]byte, OffsetRange, string]{
	MXF: {
		Bytes:         []byte{0x06, 0x0E, 0x2B, 0x34, 0x02, 0x05, 0x01, 0x01, 0x0D, 0x01, 0x02, 0x01, 0x01, 0x02},
		Offset:        OffsetRange{End: 65535},
		NameExtension: "mxf",
		Description:   "Material Exchange Format file",
		Tag:           MXF,
	},
}

var knownSignatures15 = map[FileType]HexSignature[[]byte, uint64, []string]{
	ZERO: {
		Bytes:         []byte{0x00},
		Offset:        0,
		NameExtension: []string{"PIC", "PIF", "SEA", "YTR"},
		Description:   "IBM Storyboard bitmap file, Windows Program Information File, Mac Stuffit Self-Extracting Archive, IRIS OCR data file",
		Tag:           ZERO,
	},
}
