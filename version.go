
package shwild

const (

	VersionMajor int16		=	0
	VersionMinor int16		=	1
	VersionRevision int16	=	1

	Version		int64		=	int64(VersionMajor) << 48 | int64(VersionMinor) << 32 | int64(VersionRevision) << 16
)


