package sevenzip

import (
	"fmt"
	"reflect"
	"strconv"
)

type Switches struct {
	IncludeArchives                *string `cli:"-ai"`   // -ai[r[-|0]]{@listfile|!wildcard} : Include archives
	ExcludeArchives                *string `cli:"-ax"`   // -ax[r[-|0]]{@listfile|!wildcard} : Exclude archives
	OverwriteMode                  *string `cli:"-ao"`   // -ao{a|s|t|u} : Set Overwrite mode
	DisableArchiveName             *bool   `cli:"-an"`   // -an : Disable archive_name field
	LogLevel                       *int    `cli:"-bb"`   // -bb[0-3] : Set output log level
	DisableProgress                *bool   `cli:"-bd"`   // -bd : Disable progress indicator
	OutputStream                   *string `cli:"-bs"`   // -bs{o|e|p}{0|1|2} : Set output stream for output/error/progress line
	ExecutionTimeStats             *bool   `cli:"-bt"`   // -bt : Show execution time statistics
	IncludeFilenames               *string `cli:"-i"`    // -i[r[-|0]]{@listfile|!wildcard} : Include filenames
	CompressionMethod              *string `cli:"-m"`    // -m{Parameters} : Set compression method
	OutputDirectory                *string `cli:"-o"`    // -o{Directory} : Set output directory
	Password                       *string `cli:"-p"`    // -p{Password} : Set password
	RecurseSubdirectories          *bool   `cli:"-r"`    // -r[-|0] : Recurse subdirectories
	ArchiveNameMode                *string `cli:"-sa"`   // -sa{a|e|s} : Set archive name mode
	ConsoleCharset                 *string `cli:"-scc"`  // -scc{UTF-8|WIN|DOS} : Set charset for console input/output
	ListFileCharset                *string `cli:"-scs"`  // -scs{UTF-8|UTF-16LE|UTF-16BE|WIN|DOS|{id}} : Set charset for list files
	HashFunction                   *string `cli:"-scrc"` // -scrc[CRC32|CRC64|SHA1|SHA256|*] : Set hash function for x, e, h commands
	DeleteAfterCompression         *bool   `cli:"-sdel"` // -sdel : Delete files after compression
	SendByEmail                    *bool   `cli:"-seml"` // -seml[.] : Send archive by email
	CreateSFXArchive               *string `cli:"-sfx"`  // -sfx[{name}] : Create SFX archive
	ReadFromStdin                  *string `cli:"-si"`   // -si[{name}] : Read data from stdin
	LargePagesMode                 *bool   `cli:"-slp"`  // -slp : Set large pages mode
	ShowTechInfo                   *bool   `cli:"-slt"`  // -slt : Show technical information for l (List) command
	StoreHardLinks                 *bool   `cli:"-snh"`  // -snh : Store hard links as links
	StoreSymLinks                  *bool   `cli:"-snl"`  // -snl : Store symbolic links as links
	StoreSecurityInfo              *bool   `cli:"-sni"`  // -sni : Store NT security information
	StoreNTFSStreams               *bool   `cli:"-sns"`  // -sns[-] : Store NTFS alternate streams
	WriteToStdout                  *bool   `cli:"-so"`   // -so : Write data to stdout
	DisableWildcard                *bool   `cli:"-spd"`  // -spd : Disable wildcard matching for file names
	EliminateRootFolderDuplication *bool   `cli:"-spe"`  // -spe : Eliminate duplication of root folder for extract command
	UseFullPaths                   *bool   `cli:"-spf"`  // -spf : Use fully qualified file paths
	SensitiveCaseMode              *bool   `cli:"-ssc"`  // -ssc[-] : Set sensitive case mode
	CompressSharedFiles            *bool   `cli:"-ssw"`  // -ssw : Compress shared files
	ArchiveTimestamp               *bool   `cli:"-stl"`  // -stl : Set archive timestamp from the most recently modified file
	ThreadAffinityMask             *string `cli:"-stm"`  // -stm{HexMask} : Set CPU thread affinity mask (hexadecimal number)
	ExcludeArchiveType             *string `cli:"-stx"`  // -stx{Type} : Exclude archive type
	SetArchiveType                 *string `cli:"-t"`    // -t{Type} : Set type of archive
	UpdateOptions                  *string `cli:"-u"`    // -u[-][p#][q#][r#][x#][y#][z#][!newArchiveName] : Update options
	CreateVolumes                  *string `cli:"-v"`    // -v{Size}[b|k|m|g] : Create volumes
	WorkDirectory                  *string `cli:"-w"`    // -w[{path}] : Assign work directory. Empty path means a temporary directory
	ExcludeFilenames               *string `cli:"-x"`    // -x[r[-|0]]{@listfile|!wildcard} : Exclude filenames
	AssumeYes                      *bool   `cli:"-y"`    // -y : Assume Yes on all queries
}

var DefaultLogLevel int = 3
var DefaultAssumeYes bool = true

func swV2ToArgs(sw *Switches) []string {
	args := make([]string, 0)

	t := reflect.TypeOf(sw)
	v := reflect.ValueOf(sw)

	for i := 0; i < t.Elem().NumField(); i++ {
		field := t.Elem().Field(i)
		value := v.Elem().Field(i)
		cliTag := field.Tag.Get("cli")

		if cliTag == "" {
			continue
		}

		switch value.Kind() {
		case reflect.Pointer:
			if !value.IsNil() {
				switch value.Elem().Kind() {
				case reflect.Bool:
					if value.Elem().Bool() {
						args = append(args, cliTag)
					}
				case reflect.String:
					if val := value.Elem().String(); val != "" {
						args = append(args, fmt.Sprintf("%s%s", cliTag, val))
					}
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					val := value.Elem().Int()
					args = append(args, fmt.Sprintf("%s%s", cliTag, strconv.FormatInt(val, 10)))
				}
			}
		}
	}

	return args
}

func MergeSwitches(sw ...*Switches) *Switches {
	switches := NewSwitches()
	for _, s := range sw {
		if s == nil {
			continue
		}
		if s.IncludeArchives != nil {
			switches.IncludeArchives = s.IncludeArchives
		}
		if s.ExcludeArchives != nil {
			switches.ExcludeArchives = s.ExcludeArchives
		}
		if s.OverwriteMode != nil {
			switches.OverwriteMode = s.OverwriteMode
		}
		if s.DisableArchiveName != nil {
			switches.DisableArchiveName = s.DisableArchiveName
		}
		if s.LogLevel != nil {
			switches.LogLevel = s.LogLevel
		}
		if s.DisableProgress != nil {
			switches.DisableProgress = s.DisableProgress
		}
		if s.OutputStream != nil {
			switches.OutputStream = s.OutputStream
		}
		if s.ExecutionTimeStats != nil {
			switches.ExecutionTimeStats = s.ExecutionTimeStats
		}
		if s.IncludeFilenames != nil {
			switches.IncludeFilenames = s.IncludeFilenames
		}
		if s.CompressionMethod != nil {
			switches.CompressionMethod = s.CompressionMethod
		}
		if s.OutputDirectory != nil {
			switches.OutputDirectory = s.OutputDirectory
		}
		if s.Password != nil {
			switches.Password = s.Password
		}
		if s.RecurseSubdirectories != nil {
			switches.RecurseSubdirectories = s.RecurseSubdirectories
		}
		if s.ArchiveNameMode != nil {
			switches.ArchiveNameMode = s.ArchiveNameMode
		}
		if s.ConsoleCharset != nil {
			switches.ConsoleCharset = s.ConsoleCharset
		}
		if s.ListFileCharset != nil {
			switches.ListFileCharset = s.ListFileCharset
		}
		if s.HashFunction != nil {
			switches.HashFunction = s.HashFunction
		}
		if s.DeleteAfterCompression != nil {
			switches.DeleteAfterCompression = s.DeleteAfterCompression
		}
		if s.SendByEmail != nil {
			switches.SendByEmail = s.SendByEmail
		}
		if s.CreateSFXArchive != nil {
			switches.CreateSFXArchive = s.CreateSFXArchive
		}
		if s.ReadFromStdin != nil {
			switches.ReadFromStdin = s.ReadFromStdin
		}
		if s.LargePagesMode != nil {
			switches.LargePagesMode = s.LargePagesMode
		}
		if s.ShowTechInfo != nil {
			switches.ShowTechInfo = s.ShowTechInfo
		}
		if s.StoreHardLinks != nil {
			switches.StoreHardLinks = s.StoreHardLinks
		}
		if s.StoreSymLinks != nil {
			switches.StoreSymLinks = s.StoreSymLinks
		}
		if s.StoreSecurityInfo != nil {
			switches.StoreSecurityInfo = s.StoreSecurityInfo
		}
		if s.StoreNTFSStreams != nil {
			switches.StoreNTFSStreams = s.StoreNTFSStreams
		}
		if s.WriteToStdout != nil {
			switches.WriteToStdout = s.WriteToStdout
		}
		if s.DisableWildcard != nil {
			switches.DisableWildcard = s.DisableWildcard
		}
		if s.EliminateRootFolderDuplication != nil {
			switches.EliminateRootFolderDuplication = s.EliminateRootFolderDuplication
		}
		if s.UseFullPaths != nil {
			switches.UseFullPaths = s.UseFullPaths
		}
		if s.SensitiveCaseMode != nil {
			switches.SensitiveCaseMode = s.SensitiveCaseMode
		}
		if s.CompressSharedFiles != nil {
			switches.CompressSharedFiles = s.CompressSharedFiles
		}
		if s.ArchiveTimestamp != nil {
			switches.ArchiveTimestamp = s.ArchiveTimestamp
		}
		if s.ThreadAffinityMask != nil {
			switches.ThreadAffinityMask = s.ThreadAffinityMask
		}
		if s.ExcludeArchiveType != nil {
			switches.ExcludeArchiveType = s.ExcludeArchiveType
		}
		if s.SetArchiveType != nil {
			switches.SetArchiveType = s.SetArchiveType
		}
		if s.UpdateOptions != nil {
			switches.UpdateOptions = s.UpdateOptions
		}
		if s.CreateVolumes != nil {
			switches.CreateVolumes = s.CreateVolumes
		}
		if s.WorkDirectory != nil {
			switches.WorkDirectory = s.WorkDirectory
		}
		if s.ExcludeFilenames != nil {
			switches.ExcludeFilenames = s.ExcludeFilenames
		}
		if s.AssumeYes != nil {
			switches.AssumeYes = s.AssumeYes
		}
	}

	return switches
}

func NewSwitches() *Switches {
	return &Switches{
		LogLevel:  &DefaultLogLevel,
		AssumeYes: &DefaultAssumeYes,
	}
}

func (sw *Switches) SetIncludeArchives(value string) *Switches {
	sw.IncludeArchives = &value
	return sw
}

func (sw *Switches) SetExcludeArchives(value string) *Switches {
	sw.ExcludeArchives = &value
	return sw
}

func (sw *Switches) SetOverwriteMode(value string) *Switches {
	sw.OverwriteMode = &value
	return sw
}

func (sw *Switches) SetDisableArchiveName(b bool) *Switches {
	sw.DisableArchiveName = &b
	return sw
}

func (sw *Switches) SetLogLevel(i int) *Switches {
	sw.LogLevel = &i
	return sw
}

func (sw *Switches) SetDisableProgress(b bool) *Switches {
	sw.DisableProgress = &b
	return sw
}

func (sw *Switches) SetOutputStream(value string) *Switches {
	sw.OutputStream = &value
	return sw
}

func (sw *Switches) SetExecutionTimeStats(b bool) *Switches {
	sw.ExecutionTimeStats = &b
	return sw
}

func (sw *Switches) SetIncludeFilenames(value string) *Switches {
	sw.IncludeFilenames = &value
	return sw
}

func (sw *Switches) SetCompressionMethod(value string) *Switches {
	sw.CompressionMethod = &value
	return sw
}

func (sw *Switches) SetOutputDirectory(value string) *Switches {
	sw.OutputDirectory = &value
	return sw
}

func (sw *Switches) SetPassword(value string) *Switches {
	sw.Password = &value
	return sw
}

func (sw *Switches) SetRecurseSubdirectories(b bool) *Switches {
	sw.RecurseSubdirectories = &b
	return sw
}

func (sw *Switches) SetArchiveNameMode(value string) *Switches {
	sw.ArchiveNameMode = &value
	return sw
}

func (sw *Switches) SetConsoleCharset(value string) *Switches {
	sw.ConsoleCharset = &value
	return sw
}

func (sw *Switches) SetListFileCharset(value string) *Switches {
	sw.ListFileCharset = &value
	return sw
}

func (sw *Switches) SetHashFunction(value string) *Switches {
	sw.HashFunction = &value
	return sw
}

func (sw *Switches) SetDeleteAfterCompression(b bool) *Switches {
	sw.DeleteAfterCompression = &b
	return sw
}

func (sw *Switches) SetSendByEmail(b bool) *Switches {
	sw.SendByEmail = &b
	return sw
}

func (sw *Switches) SetCreateSFXArchive(value string) *Switches {
	sw.CreateSFXArchive = &value
	return sw
}

func (sw *Switches) SetReadFromStdin(value string) *Switches {
	sw.ReadFromStdin = &value
	return sw
}

func (sw *Switches) SetLargePagesMode(b bool) *Switches {
	sw.LargePagesMode = &b
	return sw
}

func (sw *Switches) SetShowTechInfo(b bool) *Switches {
	sw.ShowTechInfo = &b
	return sw
}

func (sw *Switches) SetStoreHardLinks(b bool) *Switches {
	sw.StoreHardLinks = &b
	return sw
}

func (sw *Switches) SetStoreSymLinks(b bool) *Switches {
	sw.StoreSymLinks = &b
	return sw
}

func (sw *Switches) SetStoreSecurityInfo(b bool) *Switches {
	sw.StoreSecurityInfo = &b
	return sw
}

func (sw *Switches) SetStoreNTFSStreams(b bool) *Switches {
	sw.StoreNTFSStreams = &b
	return sw
}

func (sw *Switches) SetWriteToStdout(b bool) *Switches {
	sw.WriteToStdout = &b
	return sw
}

func (sw *Switches) SetDisableWildcard(b bool) *Switches {
	sw.DisableWildcard = &b
	return sw
}

func (sw *Switches) SetEliminateRootFolderDuplication(b bool) *Switches {
	sw.EliminateRootFolderDuplication = &b
	return sw
}

func (sw *Switches) SetUseFullPaths(b bool) *Switches {
	sw.UseFullPaths = &b
	return sw
}

func (sw *Switches) SetSensitiveCaseMode(b bool) *Switches {
	sw.SensitiveCaseMode = &b
	return sw
}

func (sw *Switches) SetCompressSharedFiles(b bool) *Switches {
	sw.CompressSharedFiles = &b
	return sw
}

func (sw *Switches) SetArchiveTimestamp(b bool) *Switches {
	sw.ArchiveTimestamp = &b
	return sw
}

func (sw *Switches) SetThreadAffinityMask(value string) *Switches {
	sw.ThreadAffinityMask = &value
	return sw
}

func (sw *Switches) SetExcludeArchiveType(value string) *Switches {
	sw.ExcludeArchiveType = &value
	return sw
}

func (sw *Switches) SetSetArchiveType(value string) *Switches {
	sw.SetArchiveType = &value
	return sw
}

func (sw *Switches) SetUpdateOptions(value string) *Switches {
	sw.UpdateOptions = &value
	return sw
}

func (sw *Switches) SetCreateVolumes(value string) *Switches {
	sw.CreateVolumes = &value
	return sw
}

func (sw *Switches) SetWorkDirectory(value string) *Switches {
	sw.WorkDirectory = &value
	return sw
}

func (sw *Switches) SetExcludeFilenames(value string) *Switches {
	sw.ExcludeFilenames = &value
	return sw
}

func (sw *Switches) SetAssumeYes(b bool) *Switches {
	sw.AssumeYes = &b
	return sw
}

func (sw *Switches) Args() []string {
	return swV2ToArgs(sw)
}
