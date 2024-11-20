package ft

import (
	"fmt"
	"regexp"
	"slices"
)

type FileTypeDescription struct {
	ShortName   string
	Description string
	Extensions  []string
}

var shortNamesIndex map[FileType]string = map[FileType]string{
	SHEBANG:              "SHEBANG",
	CLARIS_WORKS:         "CLARIS_WORKS",
	LOTUS_123_V1:         "LOTUS_123_V1",
	LOTUS_123_V3:         "LOTUS_123_V3",
	LOTUS_123_V4_V5:      "LOTUS_123_V4_V5",
	LOTUS_123_V9:         "LOTUS_123_V9",
	AMIGA_HUNK_EXE:       "AMIGA_HUNK_EXE",
	QUARK_EXPRESS:        "QUARK_EXPRESS",
	PASSWORD_GORILLA:     "PASSWORD_GORILLA",
	LIBPCAP:              "LIBPCAP",
	LIBPCAP_NS:           "LIBPCAP_NS",
	PCAPNPG:              "PCAPNPG",
	RPM:                  "RPM",
	SQLITE3:              "SQLITE3",
	AMAZON_KINDLE_UP:     "AMAZON_KINDLE_UP",
	DOOM_WAD:             "DOOM_WAD",
	ZERO:                 "ZERO",
	PALM_PILOT:           "PALM_PILOT",
	PALM_DSK_CALENDAR:    "PALM_DSK_CALENDAR",
	PALM_DSK_TODO:        "PALM_DSK_TODO",
	PALM_DSK_CALENDAR2:   "PALM_DSK_CALENDAR2",
	TELEGRAM_DSK:         "TELEGRAM_DSK",
	TELEGRAM_DSK_ENC:     "TELEGRAM_DSK_ENC",
	PALM_DSK_DATA:        "PALM_DSK_DATA",
	ICON:                 "ICON",
	APPLE_ICON_FORMAT:    "APPLE_ICON_FORMAT",
	THREE_GPP:            "THREE_GPP",
	HEIC:                 "HEIC",
	Z_LZW:                "Z_LZW",
	Z_LZH:                "Z_LZH",
	LZH0:                 "LZH0",
	LZH5:                 "LZH5",
	AMI_BACK:             "AMI_BACK",
	AMI_BACK_IDX:         "AMI_BACK_IDX",
	BPLIST:               "BPLIST",
	BZ2:                  "BZ2",
	GIF:                  "GIF",
	TIFF:                 "TIFF",
	BIG_TIFF:             "BIG_TIFF",
	CANON_RAW_V2:         "CANON_RAW_V2",
	KODAK_CIN:            "KODAK_CIN",
	RNC_V1:               "RNC_V1",
	RNC_V2:               "RNC_V2",
	NURU_IMAGE:           "NURU_IMAGE",
	NURU_PALETTE:         "NURU_PALETTE",
	SMPTE_DPX:            "SMPTE_DPX",
	OPEN_EXR:             "OPEN_EXR",
	BPG:                  "BPG",
	JPEG_RAW:             "JPEG_RAW",
	JPEG_2000:            "JPEG_2000",
	QUI:                  "QUI",
	IIF_ILBM:             "IIF_ILBM",
	IIF_VOICE:            "IIF_VOICE",
	IIF_AMIGA_CB:         "IIF_AMIGA_CB",
	IIF_ANI_BMP:          "IIF_ANI_BMP",
	IIF_ANI_CEL:          "IIF_ANI_CEL",
	IIF_FAX_IMG:          "IIF_FAX_IMG",
	IIF_FT:               "IIF_FT",
	IIF_MUS_SCORE_SIMPLE: "IIF_MUS_SCORE_SIMPLE",
	IIF_MUS_SCORE:        "IIF_MUS_SCORE",
	IIF_YUV_IMAGE:        "IIF_YUV_IMAGE",
	IIF_AMIGA_FVM:        "IIF_AMIGA_FVM",
	IIF_AIFF:             "IIF_AIFF",
	LZ:                   "LZ",
	CPIO:                 "CPIO",
	DOS_MZ:               "DOS_MZ",
	SMART_SNIFF:          "SMART_SNIFF",
	DOS_ZM:               "DOS_ZM",
	ZIP:                  "ZIP",
	RAR_V1:               "RAR_V1",
	RAR_V5:               "RAR_V5",
	ELF:                  "ELF",
	PNG:                  "PNG",
	HDF4:                 "HDF4",
	HDF5:                 "HDF5",
	COM:                  "COM",
	JAVA_CLASS:           "JAVA_CLASS",
	UTF8_TXT:             "UTF8_TXT",
	UTF16LE_TXT:          "UTF16LE_TXT",
	UTF16BE_TXT:          "UTF16BE_TXT",
	UTF32LE_TXT:          "UTF32LE_TXT",
	UTF32BE_TXT:          "UTF32BE_TXT",
	UTF7_TXT:             "UTF7_TXT",
	SCSU_TXT:             "SCSU_TXT",
	EBCDIC_TXT:           "EBCDIC_TXT",
	MACHO_BIN32:          "MACHO_BIN32",
	MACHO_BIN64:          "MACHO_BIN64",
	JKS:                  "JKS",
	MACHO_BIN32R:         "MACHO_BIN32R",
	MACHO_BIN64R:         "MACHO_BIN64R",
	PS:                   "PS",
	EPS3:                 "EPS3",
	EPS31:                "EPS31",
	CHM:                  "CHM",
	HLP:                  "HLP",
	PDF:                  "PDF",
	ASF:                  "ASF",
	MSSDI:                "MSSDI",
	OGG:                  "OGG",
	PSD:                  "PSD",
	WAV:                  "WAV",
	AVI:                  "AVI",
	MP3:                  "MP3",
	MP3V2:                "MP3V2",
	BMP:                  "BMP",
	ISO:                  "ISO",
	CDI:                  "CDI",
	MGW:                  "MGW",
	NES:                  "NES",
	D64:                  "D64",
	G64:                  "G64",
	D81:                  "D81",
	T64:                  "T64",
	CRT64:                "CRT64",
	FITS:                 "FITS",
	FLAC:                 "FLAC",
	MIDI:                 "MIDI",
	MSCOM:                "MSCOM",
	DEX:                  "DEX",
	VDMK:                 "VDMK",
	VMWARE4VDDF:          "VMWARE4VDDF",
	CRX:                  "CRX",
	FH8:                  "FH8",
	APPLE_WORKS5:         "APPLE_WORKS5",
	APPLE_WORKS6:         "APPLE_WORKS6",
	ROXIO_TOAST:          "ROXIO_TOAST",
	APPLE_DMG:            "APPLE_DMG",
	XAR:                  "XAR",
	WINDOWS_DAT:          "WINDOWS_DAT",
	NINTENDO_SYS_ROM:     "NINTENDO_SYS_ROM",
	TAR:                  "TAR",
	OAR:                  "OAR",
	TOX:                  "TOX",
	MLV:                  "MLV",
	WINDOWS_UPDATE:       "WINDOWS_UPDATE",
	SEVEN_Z:              "SEVEN_Z",
	GZIP:                 "GZIP",
	XZ:                   "XZ",
	LZ4:                  "LZ4",
	CAB:                  "CAB",
	MS_QUANTUM:           "MS_QUANTUM",
	FLIF:                 "FLIF",
	MKV:                  "MKV",
	STG:                  "STG",
	DJVU:                 "DJVU",
	DER:                  "DER",
	PEM_CRT:              "PEM_CRT",
	PEM_CSR:              "PEM_CSR",
	PEM_KEY_PKCS8:        "PEM_KEY_PKCS8",
	PEM_KEY_DSA:          "PEM_KEY_DSA",
	PEM_KEY_RSA:          "PEM_KEY_RSA",
	PPK2:                 "PPK2",
	PPK3:                 "PPK3",
	OPENSSH_PRIVK:        "OPENSSH_PRIVK",
	OPENSSH_PUBK:         "OPENSSH_PUBK",
	DCM:                  "DCM",
	WOFF:                 "WOFF",
	WOFF2:                "WOFF2",
	XML8:                 "XML8",
	XML16LE:              "XML16LE",
	XML16BE:              "XML16BE",
	XML32LE:              "XML32LE",
	XML32BE:              "XML32BE",
	XML_EBCDIC:           "XML_EBCDIC",
	WASM:                 "WASM",
	LEP:                  "LEP",
	SWF:                  "SWF",
	DEB:                  "DEB",
	WEBP:                 "WEBP",
	UBOOT:                "UBOOT",
	RTF:                  "RTF",
	MSTAPE:               "MSTAPE",
	MPEG_TS:              "MPEG_TS",
	MPEG_PS:              "MPEG_PS",
	MPEG_VS:              "MPEG_VS",
	MP4_ISO:              "MP4_ISO",
	MP4:                  "MP4",
	ZLIB_NC:              "ZLIB_NC",
	ZLIB_BS:              "ZLIB_BS",
	ZLIB_DC:              "ZLIB_DC",
	ZLIB_BC:              "ZLIB_BC",
	ZLIB_NCPD:            "ZLIB_NCPD",
	ZLIB_BSPD:            "ZLIB_BSPD",
	ZLIB_DCPD:            "ZLIB_DCPD",
	ZLIB_BCPD:            "ZLIB_BCPD",
	LZFSE:                "LZFSE",
	ORC:                  "ORC",
	AVRO:                 "AVRO",
	RC_CFF:               "RC_CFF",
	RBXL:                 "RBXL",
	PHOTOCAP_OTMPL:       "PHOTOCAP_OTMPL",
	PHOTOCAP_VEC:         "PHOTOCAP_VEC",
	PHOTOCAP_TMPL:        "PHOTOCAP_TMPL",
	PARQUET:              "PARQUET",
	EZ2:                  "EZ2",
	EZ3:                  "EZ3",
	LUA_BYTECODE:         "LUA_BYTECODE",
	MACOS_SYMLINK:        "MACOS_SYMLINK",
	MSZI:                 "MSZI",
	EML:                  "EML",
	TDE:                  "TDE",
	KDB:                  "KDB",
	PGP:                  "PGP",
	ZST:                  "ZST",
	QUICK_ZIP_RS:         "QUICK_ZIP_RS",
	SML:                  "SML",
	PEF:                  "PEF",
	SRT:                  "SRT",
	VPK:                  "VPK",
	ACE:                  "ACE",
	ARJ:                  "ARJ",
	IS_CAB:               "IS_CAB",
	KWAJ:                 "KWAJ",
	SZDD:                 "SZDD",
	ZOD:                  "ZOD",
	PBM_P1:               "PBM_P1",
	PBM_P4:               "PBM_P4",
	PBM_P2:               "PBM_P2",
	PBM_P5:               "PBM_P5",
	PBM_P3:               "PBM_P3",
	PBM_P6:               "PBM_P6",
	WMF:                  "WMF",
	XCF_GIMP:             "XCF_GIMP",
	XPM:                  "XPM",
	AFF:                  "AFF",
	ENCASE_EWF_V2:        "ENCASE_EWF_V2",
	ENCASE_EWF_V1:        "ENCASE_EWF_V1",
	QCOW:                 "QCOW",
	ANI:                  "ANI",
	CDA:                  "CDA",
	QCP:                  "QCP",
	SHOCKWAVE_DCR:        "SHOCKWAVE_DCR",
	MM_DIR:               "MM_DIR",
	FLV:                  "FLV",
	VDI:                  "VDI",
	VHD:                  "VHD",
	VHDX:                 "VHDX",
	ISZ:                  "ISZ",
	DAA:                  "DAA",
	EVT:                  "EVT",
	EVT_XML:              "EVT_XML",
	SDB:                  "SDB",
	GRP:                  "GRP",
	ICM:                  "ICM",
	MSREG_HIV:            "MSREG_HIV",
	MSOUTLOOK_PST:        "MSOUTLOOK_PST",
	DRACO:                "DRACO",
	GRIBV1V2:             "GRIBV1V2",
	BLENDER:              "BLENDER",
	JXL:                  "JXL",
	TTF:                  "TTF",
	OTF:                  "OTF",
	MODF:                 "MODF",
	MSWIM:                "MSWIM",
	SLOB:                 "SLOB",
	SJD:                  "SJD",
	CVOCF:                "CVOCF",
	AUSND:                "AUSND",
	OGL_PFB:              "OGL_PFB",
	HAZELR:               "HAZELR",
	FLP:                  "FLP",
	FLMP:                 "FLMP",
	VORENC_DPM:           "VORENC_DPM",
	MSISAM:               "MSISAM",
	MSACCDB:              "MSACCDB",
	MSMDB:                "MSMDB",
	DRW:                  "DRW",
	DSSV2:                "DSSV2",
	DSSV3:                "DSSV3",
	ADX:                  "ADX",
	ADOBE_INDD:           "ADOBE_INDD",
	MXF:                  "MXF",
	SKF:                  "SKF",
	DT2D:                 "DT2D",
	MBBTCW:               "MBBTCW",
	DESKMATE_DOC:         "DESKMATE_DOC",
	NRI:                  "NRI",
	WKS:                  "WKS",
	SIB_MUS:              "SIB_MUS",
	DSP:                  "DSP",
	AMR:                  "AMR",
	SKYPE_SILK:           "SKYPE_SILK",
	RADIANCE_HDR:         "RADIANCE_HDR",
	VBE:                  "VBE",
	CDB:                  "CDB",
	EXTM3U:               "EXTM3U",
	M2AR:                 "M2AR",
	CAPCOM_PAK:           "CAPCOM_PAK",
	CAPCOM_ARC:           "CAPCOM_ARC",
	INTERLEAF_PL:         "INTERLEAF_PL",
	NIFTI:                "NIFTI",
	NIFTI_PAIR:           "NIFTI_PAIR",
	RAF64:                "RAF64",
	VISRC:                "VISRC",
	HL7:                  "HL7",
	SAP_PWRDATA_V1:       "SAP_PWRDATA_V1",
	ARC:                  "ARC",
	PGP_APUBK:            "PGP_APUBK",
	CNT:                  "CNT",
	VDRM:                 "VDRM",
	TRID:                 "TRID",
	SHW4:                 "SHW4",
	SHW5:                 "SHW5",
	SHR5:                 "SHR5",
	SHB5:                 "SHB5",
	RMMP:                 "RMMP",
	ASTM_E57:             "ASTM_E57",
	CROWD_STRIKE_CF:      "CROWD_STRIKE_CF",
	UCAS:                 "UCAS",
	UTOC:                 "UTOC",
	COMODORE64_BIN:       "COMODORE64_BIN",
}

func NotEmptyStrings(strings []string) func(yield func(element string) bool) {
	return func(yield func(element string) bool) {
		for _, el := range strings {
			if el != "" && !yield(el) {
				return
			}
		}
	}
}

func filterEmptyExtensions[E NameExtensionPattern](pattern E) []string {
	var extensions []string
	switch value := any(pattern); value.(type) {
	case string:
		extensions = []string{value.(string)}
	case []string:
		extensions = slices.Clone(value.([]string))
	case *regexp.Regexp:
		extensions = []string{value.(*regexp.Regexp).String()}
	default:
		extensions = []string{}
	}
	result := []string{}
	for extension := range NotEmptyStrings(extensions) {
		result = append(result, extension)
	}
	return result
}

func GetFileTypeDescription(fileType FileType) (*FileTypeDescription, error) {
	hexSignature01, ok := knownSignatures01[fileType]
	if ok {
		return &FileTypeDescription{
			ShortName:   shortNamesIndex[fileType],
			Description: hexSignature01.Description,
			Extensions:  filterEmptyExtensions(hexSignature01.NameExtension),
		}, nil
	}
	hexSignature02, ok := knownSignatures02[fileType]
	if ok {
		return &FileTypeDescription{
			ShortName:   shortNamesIndex[fileType],
			Description: hexSignature02.Description,
			Extensions:  filterEmptyExtensions(hexSignature02.NameExtension),
		}, nil
	}
	hexSignature03, ok := knownSignatures03[fileType]
	if ok {
		return &FileTypeDescription{
			ShortName:   shortNamesIndex[fileType],
			Description: hexSignature03.Description,
			Extensions:  filterEmptyExtensions(hexSignature03.NameExtension),
		}, nil
	}
	hexSignature04, ok := knownSignatures04[fileType]
	if ok {
		return &FileTypeDescription{
			ShortName:   shortNamesIndex[fileType],
			Description: hexSignature04.Description,
			Extensions:  filterEmptyExtensions(hexSignature04.NameExtension),
		}, nil
	}
	hexSignature05, ok := knownSignatures05[fileType]
	if ok {
		return &FileTypeDescription{
			ShortName:   shortNamesIndex[fileType],
			Description: hexSignature05.Description,
			Extensions:  filterEmptyExtensions(hexSignature05.NameExtension),
		}, nil
	}
	hexSignature06, ok := knownSignatures06[fileType]
	if ok {
		return &FileTypeDescription{
			ShortName:   shortNamesIndex[fileType],
			Description: hexSignature06.Description,
			Extensions:  filterEmptyExtensions(hexSignature06.NameExtension),
		}, nil
	}
	hexSignature07, ok := knownSignatures07[fileType]
	if ok {
		return &FileTypeDescription{
			ShortName:   shortNamesIndex[fileType],
			Description: hexSignature07.Description,
			Extensions:  filterEmptyExtensions(hexSignature07.NameExtension),
		}, nil
	}
	hexSignature08, ok := knownSignatures08[fileType]
	if ok {
		return &FileTypeDescription{
			ShortName:   shortNamesIndex[fileType],
			Description: hexSignature08.Description,
			Extensions:  filterEmptyExtensions(hexSignature08.NameExtension),
		}, nil
	}
	hexSignature09, ok := knownSignatures09[fileType]
	if ok {
		return &FileTypeDescription{
			ShortName:   shortNamesIndex[fileType],
			Description: hexSignature09.Description,
			Extensions:  filterEmptyExtensions(hexSignature09.NameExtension),
		}, nil
	}
	hexSignature10, ok := knownSignatures10[fileType]
	if ok {
		return &FileTypeDescription{
			ShortName:   shortNamesIndex[fileType],
			Description: hexSignature10.Description,
			Extensions:  filterEmptyExtensions(hexSignature10.NameExtension),
		}, nil
	}
	hexSignature11, ok := knownSignatures11[fileType]
	if ok {
		return &FileTypeDescription{
			ShortName:   shortNamesIndex[fileType],
			Description: hexSignature11.Description,
			Extensions:  filterEmptyExtensions(hexSignature11.NameExtension),
		}, nil
	}
	hexSignature12, ok := knownSignatures12[fileType]
	if ok {
		return &FileTypeDescription{
			ShortName:   shortNamesIndex[fileType],
			Description: hexSignature12.Description,
			Extensions:  filterEmptyExtensions(hexSignature12.NameExtension),
		}, nil
	}
	hexSignature13, ok := knownSignatures13[fileType]
	if ok {
		return &FileTypeDescription{
			ShortName:   shortNamesIndex[fileType],
			Description: hexSignature13.Description,
			Extensions:  filterEmptyExtensions(hexSignature13.NameExtension),
		}, nil
	}
	hexSignature14, ok := knownSignatures14[fileType]
	if ok {
		return &FileTypeDescription{
			ShortName:   shortNamesIndex[fileType],
			Description: hexSignature14.Description,
			Extensions:  filterEmptyExtensions(hexSignature14.NameExtension),
		}, nil
	}
	hexSignature15, ok := knownSignatures15[fileType]
	if ok {
		return &FileTypeDescription{
			ShortName:   shortNamesIndex[fileType],
			Description: hexSignature15.Description,
			Extensions:  filterEmptyExtensions(hexSignature15.NameExtension),
		}, nil
	}
	return nil, fmt.Errorf("")
}
