package sdk

import "github.com/0chain/gosdk/core/sys"

const DefaultBlocksPerMarker int = 10

// DownloadOption set download option
type DownloadOption func(do *DownloadOptions)

func WithVerifyDownload(shouldVerify bool) DownloadOption {
	return func(do *DownloadOptions) {
		do.verifyDownload = shouldVerify
	}
}

func WithAllocation(obj *Allocation) DownloadOption {
	return func(do *DownloadOptions) {
		do.allocationObj = obj
	}
}

func WithBlocks(start, end int64, blocksPerMarker int) DownloadOption {
	return func(do *DownloadOptions) {
		if start > 0 && end > 0 && end >= start {
			do.isBlockDownload = true

			do.startBlock = start
			do.endBlock = end
		}

		do.blocksPerMarker = blocksPerMarker
		if do.blocksPerMarker < 1 {
			do.blocksPerMarker = DefaultBlocksPerMarker
		}

		SetNumBlockDownloads(do.blocksPerMarker)
	}
}

func WithOnlyThumbnail(thumbnail bool) DownloadOption {
	return func(do *DownloadOptions) {
		do.isThumbnailDownload = thumbnail
	}
}

func WithAuthticket(authTicket, lookupHash string) DownloadOption {
	return func(do *DownloadOptions) {
		if len(authTicket) > 0 {
			do.isViewer = true
			do.authTicket = authTicket
			do.lookupHash = lookupHash
		}
	}
}

func WithFileHandler(fileHandler sys.File) DownloadOption {
	return func(do *DownloadOptions) {
		do.fileHandler = fileHandler
		do.isFileHandlerDownload = true
	}
}
