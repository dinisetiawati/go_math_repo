package geometry

import "errors"

func CubeVolume(s int) (int, error) {
	volume := s * s * s
	if s != 0 {
		return volume, nil
	} else {
		return 0, errors.New("Zero length edge is not allowed")
	}

}
