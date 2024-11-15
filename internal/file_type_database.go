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
	MSCOM                FileType = 119 //Compound File Binary Format, a container format defined by Microsoft COM. It can contain the equivalent of files and directories. It is used by Windows Installer and for documents in older versions of Microsoft Office. It can be used by other programs as well that rely on the COM and OLE API's
	DEX                  FileType = 120 //Dalvik Executable
	VDMK                 FileType = 121 //VMDK files
	//	23 20 44 69 73 6B 20 44 65 73 63 72 69 70 74 6F	0	vmdk	VMware 4 Virtual Disk description file (split disk)
	//	43 72 32 34	0	crx	Google Chrome extension or packaged app
	//	41 47 44 33	0	fh8	FreeHand 8 document
	//	05 07 00 00 42 4F 42 4F 05 07 00 00 00 00 00 00 00 00 00 00 00 01	0	cwk	AppleWorks 5 document
	//	06 07 E1 00 42 4F 42 4F 06 07 E1 00 00 00 00 00 00 00 00 00 00 01	0	cwk AppleWorks 6 document
	//	[45 52 02 00 00 00] [8B 45 52 02 00 00 00]	0	toast	Roxio Toast disc image file
	//	6B 6F 6C 79	end–512	dmg	Apple Disk Image file
	//	78 61 72 21	0	xar	eXtensible ARchive format
	//	50 4D 4F 43 43 4D 4F 43	0	dat	Windows Files And Settings Transfer Repository[ See also USMT 3.0 (Win XP) and USMT 4.0 (Win 7) User Guides
	//	4E 45 53 1A	0	nes	Nintendo Entertainment System ROM file
	//	[75 73 74 61 72 00 30] [30 75 73 74 61 72 20 20 00]	257	tar	tar archive
	//	4F 41 52 ??	0	oar	OAR file archive format, where ?? is the format version
	//	74 6F 78 33	0	tox	Open source portable voxel file
	//	4D 4C 56 49	0	MLV	Magic Lantern Video file
	//	[44 43 4D 01 50 41 33 30] [50 41 33 30]	0		Windows Update Binary Delta Compression file
	//	37 7A BC AF 27 1C	0	7z	7-Zip File Format
	//	1F 8B	0	gz, tar.gz	GZIP compressed file
	//	FD 37 7A 58 5A 00	0	xz, tar.xz	XZ compression utility using LZMA2 compression
	//	04 22 4D 18	0	lz4	LZ4 Frame Format Remark: LZ4 block format does not offer any magic bytes
	//	4D 53 43 46	0	cab	Microsoft Cabinet file
	//	53 5A 44 44 88 F0 27 33	0	??_	Microsoft compressed file in Quantum format, used prior to Windows XP. File can be decompressed using Extract.exe or Expand.exe distributed with earlier versions of Windows. After compression, the last character of the original filename extension is replaced with an underscore, e.g. ‘Setup.exe’ becomes ‘Setup.ex_’
	//	46 4C 49 46	0	flif	Free Lossless Image Format
	//	1A 45 DF A3	0	mkv, mka, mks, mk3d, webm	Matroska media container, including WebM
	//	4D 49 4C 20	0	stg	"SEAN : Session Analysis" Training file. Also used in compatible software "Rpw : Rowperfect for Windows" and "RP3W : ROWPERFECT3 for Windows".
	//	41 54 26 54 46 4F 52 4D ?? ?? ?? ?? 44 4A 56	0	djvu, djv	DjVu document the following byte is either 55 (U) for single-page or 4D (M) for multi-page documents.
	//	30 82	0	der	DER encoded X.509 certificate
	//	2D 2D 2D 2D 2D 42 45 47 49 4E 20 43 45 52 54 49 46 49 43 41 54 45 2D 2D 2D 2D 2D	0	crt, pem	PEM encoded X.509 certificate
	//	2D 2D 2D 2D 2D 42 45 47 49 4E 20 43 45 52 54 49 46 49 43 41 54 45 20 52 45 51 55 45 53 54 2D 2D 2D 2D 2D	0	csr, pem	PEM encoded X.509 Certificate Signing Request
	//	2D 2D 2D 2D 2D 42 45 47 49 4E 20 50 52 49 56 41 54 45 20 4B 45 59 2D 2D 2D 2D 2D	0	key, pem	PEM encoded X.509 PKCS#8 private key
	//	2D 2D 2D 2D 2D 42 45 47 49 4E 20 44 53 41 20 50 52 49 56 41 54 45 20 4B 45 59 2D 2D 2D 2D 2D	0	key, pem	PEM encoded X.509 PKCS#1 DSA private key
	//	2D 2D 2D 2D 2D 42 45 47 49 4E 20 52 45 41 20 50 52 49 56 41 54 45 20 4B 45 59 2D 2D 2D 2D 2D	0	key, pem	PEM encoded X.509 PKCS#1 RSA private key
	//	50 75 54 54 59 2D 55 73 65 72 2D 4B 65 79 2D 46 69 6C 65 2D 32 3A	0	ppk	PuTTY private key file version 2
	//	50 75 54 54 59 2D 55 73 65 72 2D 4B 65 79 2D 46 69 6C 65 2D 33 3A	0	ppk	PuTTY private key file version 3
	//	2D 2D 2D 2D 2D 42 45 47 49 4E 20 4F 50 45 4E 53 53 48 20 50 52 49 56 41 54 45 20 4B 45 59 2D 2D 2D 2D 2D	0		OpenSSH private key file
	//	2D 2D 2D 2D 2D 42 45 47 49 4E 20 53 53 48 32 20 4B 45 59 2D 2D 2D 2D 2D	0	pub	OpenSSH public key file
	//	44 49 43 4D	128	dcm	DICOM Medical File Format
	//	77 4F 46 46	0	woff	WOFF File Format 1.0
	//	77 4F 46 32	0	woff2	WOFF File Format 2.0
	//	3C 3F 78 6D 6C 20	0, after BOM	xml	XML (UTF-8 or other 8-bit encodings)
	//	3C 00 3F 00 78 00 6D 00 6C 00 20	0, after BOM	xml	XML (UTF-16LE)
	//	00 3C 00 3F 00 78 00 6D 00 6C 00 20	0, after BOM	xml	XML (UTF-16BE)
	//	3C 00 00 00 3F 00 00 00 78 00 00 00 6D 00 00 00 6C 00 00 00 20 00 00 00	0, after BOM	xml	XML (UTF-32LE)
	//	00 00 00 3C 00 00 00 3F 00 00 00 78 00 00 00 6D 00 00 00 6C 00 00 00 20	0, after BOM	xml	XML (UTF-32BE)
	//	4C 6F A7 94 93 40	0, after BOM	xml	XML (EBCDIC)
	//	00 61 73 6D	0	wasm	WebAssembly binary format
	//	CF 84 01	0	lep	Lepton compressed JPEG image
	//	[43 57 53] [46 57 53]	0	swf	Adobe Flash .swf
	//	21 3C 61 72 63 68 3E 0A	0	deb	linux deb file
	//	52 49 46 46 ?? ?? ?? ?? 57 45 42 50	0	webp	Google WebP image file, where ?? ?? ?? ?? is the file size. More information on WebP File Header
	//	27 05 19 56	0		U-Boot / uImage. Das U-Boot Universal Boot Loader
	//	7B 5C 72 74 66 31	0	rtf	Rich Text Format
	//	54 41 50 45	0		Microsoft Tape Format
	//	47	0, 0xBC, 0x178, ..., (every 188th byte)	ts, tsv, tsa, mpg, mpeg	MPEG Transport Stream (MPEG-2 Part 1)
	//	00 00 01 BA	0	m2p, vob, mpg, mpeg	MPEG Program Stream (MPEG-1 Part 1 (essentially identical) and MPEG-2 Part 1)
	//	00 00 01 B3	0	mpg, mpeg	MPEG-1 video and MPEG-2 video (MPEG-1 Part 2 and MPEG-2 Part 2)
	//	66 74 79 70 69 73 6F 6D	4	mp4	ISO Base Media file (MPEG-4)
	//	66 74 79 70 4D 53 4E 56	4	mp4	MPEG-4 video file
	//	78 01	0	zlib	No Compression (no preset dictionary)
	//	78 5E	0	zlib	Best speed (no preset dictionary)
	//	78 9C	0	zlib	Default Compression (no preset dictionary)
	//	78 DA	0	zlib	Best Compression (no preset dictionary)
	//	78 20	0	zlib	No Compression (with preset dictionary)
	//	78 7D	0	zlib	Best speed (with preset dictionary)
	//	78 BB	0	zlib	Default Compression (with preset dictionary)
	//	78 F9	0	zlib	Best Compression (with preset dictionary)
	//	62 76 78 32	bvx2	0	lzfse	LZFSE - Lempel-Ziv style data compression algorithm using Finite State Entropy coding. OSS by Apple.
	//	4F 52 43	ORC	0	orc	Apache ORC (Optimized Row Columnar) file format
	//	4F 62 6A 01	Obj␁	0	avro	Apache Avro binary file format
	//	53 45 51 36	SEQ6	0	rc	RCFile columnar file format
	//	3C 72 6F 62 6C 6F 78 21	<roblox!	0	rbxl	Roblox place file
	//	65 87 78 56	e‡xV	0	p25, obt	PhotoCap Object Templates
	//	55 55 AA AA	UUªª	0	pcv	PhotoCap Vector
	//	78 56 34	xV4	0	pbt, pdt, pea, peb, pet, pgt, pict, pjt, pkt, pmt	PhotoCap Template
	//	50 41 52 31	PAR1	0		Apache Parquet columnar file format
	//	45 4D 58 32	EMX2	0	ez2	Emulator Emaxsynth samples
	//	45 4D 55 33	EMU3	0	ez3, iso	Emulator III synth samples
	//	1B 4C 75 61	␛Lua	0	luac	Lua bytecode[72]
	//	[62 6F 6F 6B 00 00 00 00] [6D 61 72 6B 00 00 00 00]	book␀␀␀␀mark␀␀␀␀	0	alias	macOS file Alias (Symbolic link)
	//	[5B 5A 6F 6E 65 54 72 61] [6E 73 66 65 72 5D]	[ZoneTransfer]	0	Identifier	Microsoft Zone Identifier for URL Security Zones
	//	52 65 63 65 69 76 65 64 3A	Received:	0	eml	Email Message var5
	//	20 02 01 62 A0 1E AB 07 02 00 00 00	␠␂␁b⍽␞«␇␂␀␀␀	0	tde	Tableau Datasource
	//	37 48 03 02 00 00 00 00 58 35 30 39 4B 45 59	7H␃␂␀␀␀␀X509KEY	0	kdb	KDB file
	//	85 ?? ?? 03	…??␃	0	pgp	PGP file
	//	28 B5 2F FD	(µ/ý	0	zst	Zstandard compress
	//	52 53 56 4B 44 41 54 41	RSVKDATA	0	rs	QuickZip rs compressed archive[79][80]
	//	3A 29 0A	:)␊	0	sml	Smile file
	//	4A 6F 79 21	Joy!	0		Preferred Executable Format
	//	31 0A 30 30	1␊00	0	srt	SubRip File
	//	34 12 AA 55	4␒ªU	0	vpk	VPK file, used to store game data for some Source Engine games
	//	2A 2A 41 43 45 2A 2A	**ACE**	7	ace	ACE (compressed file format)[citation needed]
	//	60 EA	`ê	0	arj	ARJ
	//	49 53 63 28	ISc(	0	cab	InstallShield CAB Archive File
	//	4B 57 41 4A	KWAJ	0	??_	Windows 3.1x Compressed File
	//	53 5A 44 44	SZDD	0	??_	Windows 9x Compressed File
	//	5A 4F 4F	ZOO	0	zoo	Zoo (file format)
	//	50 31 0A	P1␊	0	pbm	Portable bitmap ASCII
	//	50 34 0A	P4␊	0	pbm	Portable bitmap binary
	//	50 32 0A	P2␊	0	pgm	Portable Gray Map ASCII
	//	50 35 0A	P5␊	0	pgm	Portable Gray Map binary
	//	50 33 0A	P3␊	0	ppm	Portable Pixmap ASCII
	//	50 36 0A	P6␊	0	ppm	Portable Pixmap binary
	//	D7 CD C6 9A	×ÍÆš	0	wmf	Windows Metafile
	//	67 69 6D 70 20 78 63 66	gimp xcf	0	xcf	XCF (file format)
	//	2F 2A 20 58 50 4D 20 2A 2F	/* XPM */	0	xpm	X PixMap
	//	41 46 46	AFF	0	aff	Advanced Forensics Format
	//	45 56 46 32	EVF2	0	Ex01	EnCase EWF version 2 format
	//	45 56 46	EVF	0	e01	EnCase EWF version 1 format
	//	51 46 49	QFI	0	qcow	qcow file format
	//	52 49 46 46 ?? ?? ?? ?? 41 43 4F 4E	RIFF????ACON	0	ani	Animated cursor
	//	52 49 46 46 ?? ?? ?? ?? 43 44 44 41	RIFF????CDDA	0	cda	.cda file
	//	52 49 46 46 ?? ?? ?? ?? 51 4C 43 4D	RIFF????QLCM	0	qcp	Qualcomm PureVoice file format
	//	52 49 46 58 ?? ?? ?? ?? 46 47 44 4D (big-endian)	RIFX????FGDM	0	dcr	Adobe Shockwave[81][82][83]
	//	58 46 49 52 ?? ?? ?? ?? 4D 44 47 46 (little-endian)	XFIR????MDGF
	//	52 49 46 58 ?? ?? ?? ?? 4D 56 39 33 (big-endian)	RIFX????MV93	0	dir, dxr, drx	Macromedia Director file format[84][82][83]
	//	[58 46 49 52 ?? ?? ?? ?? 33 39 56 4D] [46 4C 56]	0	flv	Flash Video file
	//	3C 3C 3C 20 4F 72 61 63 6C 65 20 56 4D 20 56 69 72 74 75 61 6C 42 6F 78 20 44 69 73 6B 20 49 6D 61 67 65 20 3E 3E 3E	0	vdi	VirtualBox Virtual Hard Disk file format
	//	63 6F 6E 65 63 74 69 78	conectix	0	vhd	Windows Virtual PC Virtual Hard Disk file format[85]
	//	76 68 64 78 66 69 6C 65	vhdxfile	0	vhdx	Windows Virtual PC Windows 8 Virtual Hard Disk file format
	//	49 73 5A 21	IsZ!	0	isz	Compressed ISO image
	//	44 41 41	DAA	0	daa	Direct Access Archive PowerISO
	//	4C 66 4C 65	LfLe	0	evt	Windows Event Viewer file format
	//	45 6C 66 46 69 6C 65	ElfFile	0	evtx	Windows Event Viewer XML file format
	//	73 64 62 66	sdbf	8	sdb	Windows customized database
	//	50 4D 43 43	PMCC	0	grp	Windows 3.x Program Manager Program Group file format
	//	4B 43 4D 53	KCMS	0	icm	ICC profile
	//	72 65 67 66	regf	0	dat, hiv	Windows Registry file
	//	21 42 44 4E	!BDN	0	pst	Microsoft Outlook Personal Storage Table file
	//	44 52 41 43 4F	DRACO	0	drc	3D model compressed with Google Draco[86]
	//	47 52 49 42	GRIB	0	grib, grib2	Gridded data (commonly weather observations or forecasts) in the WMO GRIB or GRIB2 format[87]
	//	42 4C 45 4E 44 45 52	BLENDER	0	blend	Blender File Format[88]
	//	[00 00 00 0C 4A 58 4C 20 0D 0A 87 0A] [FF 0A]	␀␀␀␌JXL␠␍␊‡␊	0	jxl	Image encoded in the JPEG XL format
	//	00 01 00 00 00	0	ttf, tte, dfont TrueType font
	//	4F 54 54 4F	OTTO	0	otf	OpenType font
	//	23 25 4D 6F 64 75 6C 65	#%Module	0		Modulefile for Environment Modules
	//	4D 53 57 49 4D 00 00 00 D0 00 00 00 00	MSWIM␀␀␀Ð␀␀␀␀	0	wim, swm, esd	Windows Imaging Format file
	//	21 2D 31 53 4C 4F 42 1F	!-1SLOB␟	0	slob	Slob (sorted list of blobs) is a read-only, compressed data store with dictionary-like interface[92]
	//	AC ED	¬í	0		Serialized Java Data
	//	43 72 65 61 74 69 76 65 20 56 6F 69 63 65 20 46 69 6C 65 1A 1A 00	Creative Voice File	0	voc	Creative Voice file
	//	2E 73 6E 64	.snd	0	au, snd	Au audio file format
	//	DB 0A CE 00		0		OpenGL Iris Perfomer .PFB (Performer Fast Binary)[94]
	//	48 5a 4c 52 00 00 00 18	HZLR	0	hazelrules	Noodlesoft Hazel [95]
	//	46 4C 68 64	FLhd	0	flp	FL Studio Project File
	//	31 30 4C 46	10LF	0	flm	FL Studio Mobile Project File
	//	52 4b 4d 43 32 31 30	RKMC210	0		Vormetric Encryption DPM Version 2.1 Header[96]
	//	00 01 00 00 4D 53 49 53 41 4D 20 44 61 74 61 62 61 73 65	␀␁␀␀MSISAM Database	0	mny	Microsoft Money file
	//	00 01 00 00 53 74 61 6E 64 61 72 64 20 41 43 45 20 44 42	␀␁␀␀Standard ACE DB	0	accdb	Microsoft Access 2007 Database
	//	00 01 00 00 53 74 61 6E 64 61 72 64 20 4A 65 74 20 44 42	␀␁␀␀Standard Jet DB	0	mdb	Microsoft Access Database
	//	01 FF 02 04 03 02	␁ÿ␂␄␃␂	0	drw	Micrografx vector graphic file
	//	02 64 73 73	␂dss	0	dss	Digital Speech Standard (Olympus, Grundig, & Phillips) v2
	//	03 64 73 73	␃dss	0	dss	Digital Speech Standard (Olympus, Grundig, & Phillips) v3
	//	03 00 00 00 41 50 50 52	␃␀␀␀APPR	0	adx	Approach index file
	//	06 06 ED F5 D8 1D 46 E5 BD 31 EF E7 FE 74 B7 1D	␆␆íõØ␝Få½1ïçþt·␝	0	indd	Adobe InDesign document
	//	06 0E 2B 34 02 05 01 01 0D 01 02 01 01 02	␆␎+4␂␅␁␁␍␁␂␁␁␂	0-65535 (run-in)	mxf	Material Exchange Format file
	//	07 53 4B 46	␇SKF	0	skf	SkinCrafter skin file
	//	07 64 74 32 64 64 74 64	␇dt2ddtd	0	dtd	DesignTools 2D Design file
	//	0A 16 6F 72 67 2E 62 69 74 63 6F 69 6E 2E 70 72	␊␖org.bitcoin.pr	0	wallet	MultiBit Bitcoin wallet file
	//	0D 44 4F 43	␍DOC	0	doc	DeskMate Document file
	//	0E 4E 65 72 6F 49 53 4F	␎NeroISO	0	nri	Nero CD Compilation
	//	0E 57 4B 53	␎WKS	0	wks	DeskMate Worksheet
	//	0F 53 49 42 45 4C 49 55 53	␏SIBELIUS	0	sib	Sibelius Music - Score file
	//	23 20 4D 69 63 72 6F 73 6F 66 74 20 44 65 76 65 6C 6F 70 65 72 20 53 74 75 64 69 6F	# Microsoft Developer Studio	0	dsp	Microsoft Developer Studio project file
	//	23 21 41 4D 52	#!AMR	0	amr	Adaptive Multi-Rate ACELP (Algebraic Code Excited Linear Prediction) Codec, commonly audio format with GSM cell phones.
	//	23 21 53 49 4C 4B 0A	#!SILK␊	0	sil	Audio compression format developed by Skype
	//	23 3F 52 41 44 49 41 4E 43 45 0A	#?RADIANCE␊	0	hdr	Radiance High Dynamic Range image file
	//	23 40 7E 5E	#@~^	0	vbe	VBScript Encoded script
	//	0D F0 1D C0	␍ð�À	0	cdb	MikroTik WinBox Connection Database (Address Book)
	//	23 45 58 54 4D 33 55	#EXTM3U	0	m3u, m3u8 Multimedia playlist
	//	6D 64 66 00	mdf␀	0	m	M2 Archive, used by game developer M2
	//	4B 50 4B 41	KPKA	0	pak	Capcom RE Engine game data archives
	//	41 52 43	ARC	0	arc	Capcom MT Framework game data archives
	//	D0 4F 50 53	ÐOPS	0	pl	Interleaf PrinterLeaf / WorldView document format (now Broadvision QuickSilver)
	//	6E 2B 31 00	n+1	344	nii	Single file NIfTI format, used extensively in biomedical imaging.
	//	6E 69 31 00	ni1	344	hdr	Header file of a .hdr/.img pair in NIfTI format, used extensively in biomedical imaging.
	//	52 41 46 36 34	RAF64	0		Report Builder file from Digital Metaphors
	//	56 49 53 33	VIS3	0		Resource file Visionaire 3.x Engine
	//	4D 53 48 7C 42 53 48 7C	MSH|BSH|	0	hl7	Health Level Seven (HL7) Standard for electronic data exchange [1]
	//	70 77 72 64 61 74 61	pwrdata	0	pwrdata	SAP Power Monitor (version 1.1.0 and higher) data file
	//	1a 08	..	0	arc	ARC archive file
	//	2d 2d 2d 2d 2d 42 45 47 49 4e 20 50 47 50 20 50 55 42 4c 49 43 20 4b 45 49 20 42 4c 4f 43 4b 2d 2d 2d 2d 2d	-----BEGIN PGP PUBLIC KEY BLOCK-----	0	asc	Armored PGP public key
	//	3a 42 61 73 65 20	:Base	0	cnt	Windows 3.x - Windows 95 Help Contents
	//	52 49 46 46 ?? ?? ?? ?? 56 44 52 4d	RIFF????VDRM	0	vdr	VirtualDub
	//	52 59 46 46 ?? ?? ?? ?? 54 52 49 44	RIFF????TRID	0	trd	TrID
	//	52 49 46 46 ?? ?? ?? ?? 73 68 77 34	RIFF????shw4	0	shw	Corel SHOW! 4.0
	//	52 49 46 46 ?? ?? ?? ?? 73 68 77 35	RIFF????shw5	0	shw	Corel SHOW! 5.0
	//	52 49 46 46 ?? ?? ?? ?? 73 68 72 35	RIFF????shr5	0	shr	Corel SHOW! 5.0 player
	//	52 49 46 46 ?? ?? ?? ?? 73 68 62 35	RIFF????shb5	0	shb	Corel SHOW! 5.0 background
	//	52 49 46 46 ?? ?? ?? ?? 52 4d 4d 50	RIFF????RMMP	0	mmm	MacroMind Multimedia Movie or Microsoft Multimedia Movie
	//	41 53 54 4d 2d 45 35 37	ASTM-E57	0	e57	ASTM E57 3D file format
	//	aa aa aa aa	ªªªª	0	sys	Crowdstrike Channel File
	//	8C 0A 00	Œ..	0	ucas	Unreal Engine Compressed Asset Storage file
	//	2D 3D 3D 2D 2D 3D 3D 2D 2D 3D 3D 2D 2D 3D 3D 2D	-==--==--==--==-	0	utoc	Unreal Engine Table of Contents file
	//	43 36 34 46 69 6C 65 00	C64File	0	P00 (P01, P02,...)	Commodore 64 binary file (source: *.P00 format for Power64 emulator)
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
