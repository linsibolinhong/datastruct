package ringbuf

import (
	"fmt"
)

type RingBuffer struct {
	buf []byte
	begin int
	end int

	isEmpty bool
}

func New(size uint) *RingBuffer {
	if size == 0 {
		return nil
	}
	buf := &RingBuffer{}
	buf.buf = make([]byte, size)
	buf.begin = 0
	buf.end = 0
	buf.isEmpty = true
	return buf
}

func (r *RingBuffer) Empty() bool {
	return r.isEmpty
}

func (r *RingBuffer) Size() int {
	return len(r.buf)
}

func (r *RingBuffer) Begin() int {
	return r.begin
}

func (r *RingBuffer) End() int {
	return r.end
}

func (r *RingBuffer) UsedSize() int {
	if r.isEmpty {
		return 0
	}

	if r.end > r.begin {
		return r.end - r.begin
	} else {
		return r.Size() - r.begin + r.end
	}
}

func (r *RingBuffer) FreeSize() int {
	return r.Size() - r.UsedSize()
}

func (r *RingBuffer) Write(data []byte) int {
	n := len(data)

	if n == 0 {
		return 0
	}

	if n > r.FreeSize() {
		return -1
	}

	if n <= r.Size() - r.End() {
		copy(r.buf[r.End():r.End() + n], data)
	} else {
		copy(r.buf[r.End():],data[:r.Size() - r.End()])
		copy(r.buf[:n - r.Size() + r.End()], data[r.Size() - r.End():])
	}
	r.end = (r.end + n) % r.Size()
	r.isEmpty = false
	return n
}

func (r *RingBuffer) ReadData(data []byte) int {
	n := r.ReadDataAt(data, r.Begin())
	if n > 0 {
		if n == r.UsedSize() {
			r.isEmpty = true
		}
		r.begin = (r.begin + n) % r.Size()
	}
	return n
}

func (r *RingBuffer) WriteAt(data []byte, idx int) int {
	if idx < 0 || idx >= r.Size() {
		return -1
	}

	if len(data) > r.FreeSize() {
		return -1
	}

	if len(data) == 0 {
		return 0
	}

	if r.Begin() < r.End() {
		if idx < r.Begin() || idx > r.End() {
			return -1
		}

		r.end = idx
		return r.Write(data)
	} else {
		if idx > r.End() && idx < r.Begin() {
			return -1
		}
		r.end = idx
		return r.Write(data)
	}
}

func (r *RingBuffer) ReadDataAt(data []byte, idx int) int {
	if idx < 0 || idx >= r.Size() {
		return -1
	}

	n := len(data)
	if r.Begin() < r.End() {
		if idx >= r.End() || idx < r.Begin() {
			return -1
		}
		if r.UsedSize() == 0 {
			return 0
		}
		usedSize := r.End() - idx
		if n <= usedSize {
			copy(data, r.buf[idx:idx+ n])
			return n
		} else {
			copy(data, r.buf[idx:r.End()])
			return usedSize
		}
	} else {
		if idx >= r.End() && idx < r.Begin() {
			return -1
		}
		if r.UsedSize() == 0 {
			return 0
		}
		if idx >= r.Begin() {
			if n <= r.Size() - idx {
				copy(data, r.buf[idx:idx + n])
				return n
			} else {
				copy(data, r.buf[idx:r.Size()])
				copied := r.Size() - idx
				if n - copied <= r.End() {
					copy(data[copied:], r.buf[:n-copied])
					return n
				} else {
					copy(data[copied:], r.buf[:r.End()])
					return copied + r.End()
				}
			}
		} else {
			if n <= r.End() - idx {
				copy(data, r.buf[idx:idx+ n])
				return n
			} else {
				copy(data, r.buf[idx:r.End()])
				return r.End() - idx
			}
		}
	}
}

func (r *RingBuffer) ReadAt(n, idx int) ([]byte, int) {
	ret := make([]byte, n)
	nt := r.ReadDataAt(ret, idx)
	if nt >= 0 {
		ret = ret[:nt]
	}
	return ret, n
}

func (r *RingBuffer) Read(n int) ([]byte, int) {
	ret := make([]byte, n)
	nt := r.ReadData(ret)
	if nt >= 0 {
		ret = ret[:nt]
	}
	return ret, n
}

func (r *RingBuffer) Detail() string {
	udata, _ := r.ReadAt(r.UsedSize(), r.begin)
	return fmt.Sprintf("size:%d\nbegin:%d\nend:%d\nused:%d\nfree:%d\ndata:%s\n",
		r.Size(),r.Begin(), r.End(), r.UsedSize(), r.FreeSize(), string(udata))
}

func (r *RingBuffer) ReadDataWithOffset(data []byte, offset int) int {
	idx := (r.begin + offset) % r.Size()
	return r.ReadDataAt(data, idx)
}

func (r *RingBuffer) ReadWithOffset(n, offset int) ([]byte, int) {
	ret := make([]byte, n)
	nt := r.ReadDataWithOffset(ret, offset)
	if nt >= 0 {
		ret = ret[:nt]
	}
	return ret, n
}

func (r *RingBuffer) RemoveAt(idx, n int) (ret int) {
	if r.isEmpty {
		return 0
	}

	if n == 0 {
		return 0
	}

	defer func() {
		if ret > 0 {
			r.isEmpty = r.begin == r.end
		}
	}()

	if idx < 0 || idx >= r.Size() {
		return -1
	}

	if r.UsedSize() < n {
		return -1
	}

	if r.Begin() < r.End() {
		if idx >= r.End() || idx < r.Begin() {
			return -1
		}

		if r.End() - idx < n {
			return -1
		}

		copy(r.buf[idx:], r.buf[idx+n:r.End()])
		r.end = r.end - n
		return n
	} else {
		if idx >= r.End() && idx < r.Begin() {
			return -1
		}

		if idx >= r.Begin() {
			if r.End() + r.Size() - idx < n {
				return -1
			}

			if n + idx < r.Size() {
				copy(r.buf[idx:], r.buf[idx+n:])
				if n < r.End() {
					copy(r.buf[:], r.buf[n:r.End()])
					r.end -= n
				} else {
					r.end = r.End() + r.Size() - idx + n
				}
				return n
			} else {
				copy(r.buf[idx:], r.buf[(idx+n)%r.Size():r.End()])
				r.end -= n
				if r.end < 0 {
					r.end += r.Size()
				}
				return n
			}
		} else {
			if n <= r.End() - idx {
				copy(r.buf[idx:], r.buf[idx+n:r.End()])
				r.end -= n
				return n
			} else {
				return -1
			}
		}
	}
}

func (r *RingBuffer) Remove(offset, n int) int {
	if n < 0 || n >= r.Size() {
		return -1
	}
	return r.RemoveAt((r.begin + offset) % r.Size(), n)
}

func (r *RingBuffer) Clear() {
	r.begin = 0
	r.end = 0
	r.isEmpty = true
}


