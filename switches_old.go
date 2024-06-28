package sevenzip

// import (
// 	"fmt"
// 	"reflect"
// )

// type Switches struct {
// 	IncludeArchives                stringSwitches // -ai[r[-|0]]{@listfile|!wildcard} : Include archives
// 	ExcludeArchives                stringSwitches // -ax[r[-|0]]{@listfile|!wildcard} : Exclude archives
// 	OverwriteMode                  stringSwitches // -ao{a|s|t|u} : Set Overwrite mode
// 	DisableArchiveName             boolSwitches   // -an : Disable archive_name field
// 	LogLevel                       intSwitches    // -bb[0-3] : Set output log level
// 	DisableProgress                boolSwitches   // -bd : Disable progress indicator
// 	OutputStream                   stringSwitches // -bs{o|e|p}{0|1|2} : Set output stream for output/error/progress line
// 	ExecutionTimeStats             boolSwitches   // -bt : Show execution time statistics
// 	IncludeFilenames               stringSwitches // -i[r[-|0]]{@listfile|!wildcard} : Include filenames
// 	CompressionMethod              stringSwitches // -m{Parameters} : Set compression method
// 	OutputDirectory                stringSwitches // -o{Directory} : Set output directory
// 	Password                       stringSwitches // -p{Password} : Set password
// 	RecurseSubdirectories          boolSwitches   // -r[-|0] : Recurse subdirectories
// 	ArchiveNameMode                stringSwitches // -sa{a|e|s} : Set archive name mode
// 	ConsoleCharset                 stringSwitches // -scc{UTF-8|WIN|DOS} : Set charset for console input/output
// 	ListFileCharset                stringSwitches // -scs{UTF-8|UTF-16LE|UTF-16BE|WIN|DOS|{id}} : Set charset for list files
// 	HashFunction                   stringSwitches // -scrc[CRC32|CRC64|SHA1|SHA256|*] : Set hash function for x, e, h commands
// 	DeleteAfterCompression         boolSwitches   // -sdel : Delete files after compression
// 	SendByEmail                    boolSwitches   // -seml[.] : Send archive by email
// 	CreateSFXArchive               stringSwitches // -sfx[{name}] : Create SFX archive
// 	ReadFromStdin                  stringSwitches // -si[{name}] : Read data from stdin
// 	LargePagesMode                 boolSwitches   // -slp : Set large pages mode
// 	ShowTechInfo                   boolSwitches   // -slt : Show technical information for l (List) command
// 	StoreHardLinks                 boolSwitches   // -snh : Store hard links as links
// 	StoreSymLinks                  boolSwitches   // -snl : Store symbolic links as links
// 	StoreSecurityInfo              boolSwitches   // -sni : Store NT security information
// 	StoreNTFSStreams               boolSwitches   // -sns[-] : Store NTFS alternate streams
// 	WriteToStdout                  boolSwitches   // -so : Write data to stdout
// 	DisableWildcard                boolSwitches   // -spd : Disable wildcard matching for file names
// 	EliminateRootFolderDuplication boolSwitches   // -spe : Eliminate duplication of root folder for extract command
// 	UseFullPaths                   boolSwitches   // -spf : Use fully qualified file paths
// 	SensitiveCaseMode              boolSwitches   // -ssc[-] : Set sensitive case mode
// 	CompressSharedFiles            boolSwitches   // -ssw : Compress shared files
// 	ArchiveTimestamp               boolSwitches   // -stl : Set archive timestamp from the most recently modified file
// 	ThreadAffinityMask             stringSwitches // -stm{HexMask} : Set CPU thread affinity mask (hexadecimal number)
// 	ExcludeArchiveType             stringSwitches // -stx{Type} : Exclude archive type
// 	SetArchiveType                 stringSwitches // -t{Type} : Set type of archive
// 	UpdateOptions                  stringSwitches // -u[-][p#][q#][r#][x#][y#][z#][!newArchiveName] : Update options
// 	CreateVolumes                  stringSwitches // -v{Size}[b|k|m|g] : Create volumes
// 	WorkDirectory                  stringSwitches // -w[{path}] : Assign work directory. Empty path means a temporary directory
// 	ExcludeFilenames               stringSwitches // -x[r[-|0]]{@listfile|!wildcard} : Exclude filenames
// 	AssumeYes                      boolSwitches   // -y : Assume Yes on all queries
// }

// func NewSwitches() *Switches {
// 	return &Switches{
// 		IncludeArchives:                stringSwitches{switches: "-ai"},
// 		ExcludeArchives:                stringSwitches{switches: "-ax"},
// 		OverwriteMode:                  stringSwitches{switches: "-ao"},
// 		DisableArchiveName:             boolSwitches{switches: "-an"},
// 		LogLevel:                       intSwitches{switches: "-bb"},
// 		DisableProgress:                boolSwitches{switches: "-bd"},
// 		OutputStream:                   stringSwitches{switches: "-bs"},
// 		ExecutionTimeStats:             boolSwitches{switches: "-bt"},
// 		IncludeFilenames:               stringSwitches{switches: "-i"},
// 		CompressionMethod:              stringSwitches{switches: "-m"},
// 		OutputDirectory:                stringSwitches{switches: "-o"},
// 		Password:                       stringSwitches{switches: "-p"},
// 		RecurseSubdirectories:          boolSwitches{switches: "-r"},
// 		ArchiveNameMode:                stringSwitches{switches: "-sa"},
// 		ConsoleCharset:                 stringSwitches{switches: "-scc"},
// 		ListFileCharset:                stringSwitches{switches: "-scs"},
// 		HashFunction:                   stringSwitches{switches: "-scrc"},
// 		DeleteAfterCompression:         boolSwitches{switches: "-sdel"},
// 		SendByEmail:                    boolSwitches{switches: "-seml"},
// 		CreateSFXArchive:               stringSwitches{switches: "-sfx"},
// 		ReadFromStdin:                  stringSwitches{switches: "-si"},
// 		LargePagesMode:                 boolSwitches{switches: "-slp"},
// 		ShowTechInfo:                   boolSwitches{switches: "-slt"},
// 		StoreHardLinks:                 boolSwitches{switches: "-snh"},
// 		StoreSymLinks:                  boolSwitches{switches: "-snl"},
// 		StoreSecurityInfo:              boolSwitches{switches: "-sni"},
// 		StoreNTFSStreams:               boolSwitches{switches: "-sns"},
// 		WriteToStdout:                  boolSwitches{switches: "-so"},
// 		DisableWildcard:                boolSwitches{switches: "-spd"},
// 		EliminateRootFolderDuplication: boolSwitches{switches: "-spe"},
// 		UseFullPaths:                   boolSwitches{switches: "-spf"},
// 		SensitiveCaseMode:              boolSwitches{switches: "-ssc"},
// 		CompressSharedFiles:            boolSwitches{switches: "-ssw"},
// 		ArchiveTimestamp:               boolSwitches{switches: "-stl"},
// 		ThreadAffinityMask:             stringSwitches{switches: "-stm"},
// 		ExcludeArchiveType:             stringSwitches{switches: "-stx"},
// 		SetArchiveType:                 stringSwitches{switches: "-t"},
// 		UpdateOptions:                  stringSwitches{switches: "-u"},
// 		CreateVolumes:                  stringSwitches{switches: "-v"},
// 		WorkDirectory:                  stringSwitches{switches: "-w"},
// 		ExcludeFilenames:               stringSwitches{switches: "-x"},
// 		AssumeYes:                      boolSwitches{switches: "-y"},
// 	}
// }

// func (sw *Switches) Args() []string {
// 	return swToArgs(sw)
// }

// type argParser interface {
// 	Parse() []string //  Used in the cmd call
// }

// type stringSwitches struct {
// 	switches string
// 	value    string
// }

// func (so stringSwitches) Parse() []string {
// 	args := make([]string, 0)
// 	if so.value == "" {
// 		return args
// 	}
// 	args = append(args, so.switches)
// 	args = append(args, so.value)
// 	return args
// }

// func (so *stringSwitches) Set(value string) {
// 	so.value = value
// }

// func (so *stringSwitches) Unset() {
// 	so.value = ""
// }

// type intSwitches struct {
// 	switches string
// 	value    int
// 	isSet    bool
// }

// func (io intSwitches) Parse() []string {
// 	args := make([]string, 0)
// 	if !io.isSet {
// 		return args
// 	}
// 	args = append(args, io.switches)
// 	args = append(args, fmt.Sprintf("%d", io.value))
// 	return args
// }

// func (io *intSwitches) Set(value int) {
// 	io.isSet = true
// 	io.value = value
// }

// func (io *intSwitches) Unset() {
// 	io.isSet = false
// }

// type boolSwitches struct {
// 	switches string
// 	value    bool
// }

// func (bo boolSwitches) Parse() []string {
// 	if bo.value {
// 		return []string{bo.switches}
// 	}
// 	return []string{}
// }

// func (bo *boolSwitches) Set(value bool) {
// 	bo.value = value
// }

// func (bo *boolSwitches) Unset() {
// 	bo.value = false
// }

// func swToArgs(sw *Switches) []string {
// 	args := make([]string, 0)
// 	rv := reflect.Indirect(reflect.ValueOf(sw))
// 	if rv.Kind() != reflect.Struct {
// 		return args
// 	}
// 	for i := 0; i < rv.NumField(); i++ {
// 		prsr, ok := rv.Field(i).Interface().(argParser)
// 		if ok {
// 			s := prsr.Parse()
// 			if len(s) > 0 {
// 				args = append(args, s...)
// 			}
// 		}
// 	}
// 	return args
// }

// func MergeSwitches(sw ...*Switches) *Switches {
// 	switches := NewSwitches()
// 	for _, s := range sw {
// 		if s == nil {
// 			continue
// 		}
// 		if s.IncludeArchives.value != "" {
// 			switches.IncludeArchives.value = s.IncludeArchives.value
// 		}
// 		if s.ExcludeArchives.value != "" {
// 			switches.ExcludeArchives.value = s.ExcludeArchives.value
// 		}
// 		if s.OverwriteMode.value != "" {
// 			switches.OverwriteMode.value = s.OverwriteMode.value
// 		}
// 		if s.DisableArchiveName.value {
// 			switches.DisableArchiveName.value = s.DisableArchiveName.value
// 		}
// 		if s.LogLevel.value != 0 {
// 			switches.LogLevel.value = s.LogLevel.value
// 		}
// 		if s.DisableProgress.value {
// 			switches.DisableProgress.value = s.DisableProgress.value
// 		}
// 		if s.OutputStream.value != "" {
// 			switches.OutputStream.value = s.OutputStream.value
// 		}
// 		if s.ExecutionTimeStats.value {
// 			switches.ExecutionTimeStats.value = s.ExecutionTimeStats.value
// 		}
// 		if s.IncludeFilenames.value != "" {
// 			switches.IncludeFilenames.value = s.IncludeFilenames.value
// 		}
// 		if s.CompressionMethod.value != "" {
// 			switches.CompressionMethod.value = s.CompressionMethod.value
// 		}
// 		if s.OutputDirectory.value != "" {
// 			switches.OutputDirectory.value = s.OutputDirectory.value
// 		}
// 		if s.Password.value != "" {
// 			switches.Password.value = s.Password.value
// 		}
// 		if s.RecurseSubdirectories.value {
// 			switches.RecurseSubdirectories.value = s.RecurseSubdirectories.value
// 		}
// 		if s.ArchiveNameMode.value != "" {
// 			switches.ArchiveNameMode.value = s.ArchiveNameMode.value
// 		}
// 		if s.ConsoleCharset.value != "" {
// 			switches.ConsoleCharset.value = s.ConsoleCharset.value
// 		}
// 		if s.ListFileCharset.value != "" {
// 			switches.ListFileCharset.value = s.ListFileCharset.value
// 		}
// 		if s.HashFunction.value != "" {
// 			switches.HashFunction.value = s.HashFunction.value
// 		}
// 		if s.DeleteAfterCompression.value {
// 			switches.DeleteAfterCompression.value = s.DeleteAfterCompression.value
// 		}
// 		if s.SendByEmail.value {
// 			switches.SendByEmail.value = s.SendByEmail.value
// 		}
// 		if s.CreateSFXArchive.value != "" {
// 			switches.CreateSFXArchive.value = s.CreateSFXArchive.value
// 		}
// 		if s.ReadFromStdin.value != "" {
// 			switches.ReadFromStdin.value = s.ReadFromStdin.value
// 		}
// 		if s.LargePagesMode.value {
// 			switches.LargePagesMode.value = s.LargePagesMode.value
// 		}
// 		if s.ShowTechInfo.value {
// 			switches.ShowTechInfo.value = s.ShowTechInfo.value
// 		}
// 		if s.StoreHardLinks.value {
// 			switches.StoreHardLinks.value = s.StoreHardLinks.value
// 		}
// 		if s.StoreSymLinks.value {
// 			switches.StoreSymLinks.value = s.StoreSymLinks.value
// 		}
// 		if s.StoreSecurityInfo.value {
// 			switches.StoreSecurityInfo.value = s.StoreSecurityInfo.value
// 		}
// 		if s.StoreNTFSStreams.value {
// 			switches.StoreNTFSStreams.value = s.StoreNTFSStreams.value
// 		}
// 		if s.WriteToStdout.value {
// 			switches.WriteToStdout.value = s.WriteToStdout.value
// 		}
// 		if s.DisableWildcard.value {
// 			switches.DisableWildcard.value = s.DisableWildcard.value
// 		}
// 		if s.EliminateRootFolderDuplication.value {
// 			switches.EliminateRootFolderDuplication.value = s.EliminateRootFolderDuplication.value
// 		}
// 		if s.UseFullPaths.value {
// 			switches.UseFullPaths.value = s.UseFullPaths.value
// 		}
// 		if s.SensitiveCaseMode.value {
// 			switches.SensitiveCaseMode.value = s.SensitiveCaseMode.value
// 		}
// 		if s.CompressSharedFiles.value {
// 			switches.CompressSharedFiles.value = s.CompressSharedFiles.value
// 		}
// 		if s.ArchiveTimestamp.value {
// 			switches.ArchiveTimestamp.value = s.ArchiveTimestamp.value
// 		}
// 		if s.ThreadAffinityMask.value != "" {
// 			switches.ThreadAffinityMask.value = s.ThreadAffinityMask.value
// 		}
// 		if s.ExcludeArchiveType.value != "" {
// 			switches.ExcludeArchiveType.value = s.ExcludeArchiveType.value
// 		}
// 		if s.SetArchiveType.value != "" {
// 			switches.SetArchiveType.value = s.SetArchiveType.value
// 		}
// 		if s.UpdateOptions.value != "" {
// 			switches.UpdateOptions.value = s.UpdateOptions.value
// 		}
// 		if s.CreateVolumes.value != "" {
// 			switches.CreateVolumes.value = s.CreateVolumes.value
// 		}
// 		if s.WorkDirectory.value != "" {
// 			switches.WorkDirectory.value = s.WorkDirectory.value
// 		}
// 		if s.ExcludeFilenames.value != "" {
// 			switches.ExcludeFilenames.value = s.ExcludeFilenames.value
// 		}
// 		if s.AssumeYes.value {
// 			switches.AssumeYes.value = s.AssumeYes.value
// 		}
// 	}

// 	return switches
// }
