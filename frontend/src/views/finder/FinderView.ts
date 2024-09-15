import axios from '@/utils/request'
import type { AxiosProgressEvent } from 'axios'

export interface FileInfo {
  name: string;
  path: string;
  size: number;
  isDir: boolean;
  modTime: Date;
}

export interface Storage {
  title: string;
  locationType: string;
  args?: string;
  id: string;
  createdAt: Date;
  updatedAt: Date;
}

export const listStorages = async(): Promise<Storage[]> => {
  return axios.get(`/storage/list`)
}

export const listDirectory = async (id: string, path: string): Promise<FileInfo[]> => {
  return axios.get(`/storage/${id}/folder?path=${path}`)
}

export const getSP = async (id: string): Promise<FileInfo[]> => {
  return axios.get(`/storage/${id}/sp`)
}

export const fileOriDirExists = async (id: string, path: string): Promise<boolean> => {
  return axios.get(`/storage/${id}/exists?path=${path}`)
}

export const onUploadFile = async (id: string, option: any, targetPath: string, mode: 'overwrite' | 'ignore') => {
  const { fileItem, name, onProgress, onError, onSuccess } = option
  const formData = new FormData()
  const n = name || fileItem.name
  const tp = joinPaths(targetPath, getDirectoryFromRelativePath(fileItem.file.webkitRelativePath))
  formData.append('fileName', n)
  formData.append('file', fileItem.file)
  formData.append('targetPath', tp)
  formData.append('mode', mode)

  if (mode === 'ignore' && await fileOriDirExists(id, joinPaths(tp, n))) {
    option.onProgress(100)
    return
  }

  await axios.post(`/storage/${id}/upload`, formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    onUploadProgress(progressEvent: AxiosProgressEvent) {
      const percent = Math.round((progressEvent.loaded * 100) / (progressEvent.total || 999999))
      onProgress(percent)
    }
  })
    .then((res: any) => {
      onSuccess(res.data)
    })
    .catch((err) => {
      onError(err)
    })
}

function joinPaths(...segments: string[]): string {
  return segments
    .map((segment) => segment.replace(/\/+$/, '')) // 去掉末尾的斜杠
    .join('/')
    .replace(/\/{2,}/g, '/') // 防止多个斜杠
}

function getDirectoryFromRelativePath(relativePath: string): string {
  return relativePath.replace(/\/[^/]*$/, '') // 移除最后的文件名部分
}

export const getFileIcon = (file: FileInfo): string => {
  if (file.isDir) {
    return 'internal://icon-finder-folder'
  }
  switch (getFullFileExtension(file.name).toLowerCase()) {
    case '.ts':
    case '.vue':
    case '.html':
    case '.js':
    case '.css':
    case '.go':
      return 'internal://icon-finder-code'
    case '.jpg':
    case '.svg':
    case '.png':
      return 'internal://icon-finder-image'
    case '.7z':
    case '.tar.gz':
    case '.tgz':
    case '.zip':
      return 'internal://icon-finder-ys'
    case '.ai':
      return 'internal://icon-finder-ai'
    case '.pdf':
      return 'internal://icon-finder-pdf'
    case '.doc':
    case '.docx':
      return 'internal://icon-finder-doc'
    case '.ppt':
    case '.pptx':
      return 'internal://icon-finder-ppt'
    case '.txt':
    case '.md':
    case '.sum':
    case '.mod':
      return 'internal://icon-finder-txt'
    case '.csv':
    case '.xls':
    case '.xlsx':
    case '.xlsb':
    case '.xlsm':
    case '.elx':
      return 'internal://icon-finder-elx'
    case '.vsd':
      return 'internal://icon-finder-vsd'
    case '.xmap':
      return 'internal://icon-finder-xmap'
    case '.psd':
      return 'internal://icon-finder-psd'
    case '.dwg':
      return 'internal://icon-finder-dwg'
    case '.mpp':
      return 'internal://icon-finder-mpp'
    case '.mp3':
    case '.wav':
    case '.aac':
    case '.flac':
    case '.ogg':
    case '.m4a':
    case '.aiff':
    case '.alac':
    case '.wma':
    case '.opus':
      return 'internal://icon-finder-audio'
    case '.mp4':
    case '.avi':
    case '.mkv':
    case '.mov':
    case '.wmv':
    case '.flv':
    case '.webm':
    case '.mpg':
    case '.mpeg':
    case '.3gp':
    case '.hevc':
      return 'internal://icon-finder-video'
    default:
      return 'internal://icon-finder-unknown'
  }
}

function getFullFileExtension(filename: string): string {
  if (!filename) {
    return ''
  }

  // 从文件名中提取出后缀
  const lastDotIndex = filename.lastIndexOf('.')
  if (lastDotIndex === -1 || lastDotIndex === filename.length - 1) {
    return ''
  }

  // 从最后一个点开始获取文件扩展名
  let extension = filename.substring(lastDotIndex + 1)
  const secondLastDotIndex = filename.lastIndexOf('.', lastDotIndex - 1)

  // 如果有两个点，返回完整扩展名
  if (secondLastDotIndex !== -1) {
    extension = filename.substring(secondLastDotIndex + 1)
  }

  return '.' + extension
}

