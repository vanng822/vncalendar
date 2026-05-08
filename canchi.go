package vncalendar

import (
	"fmt"
	"math"
	"strings"
)

var (
	can    = []string{"Giáp", "Ất", "Bính", "Đinh", "Mậu", "Kỷ", "Canh", "Tân", "Nhâm", "Quý"}
	chi    = []string{"Tý", "Sửu", "Dần", "Mão", "Thìn", "Tị", "Ngọ", "Mùi", "Thân", "Dậu", "Tuất", "Hợi"}
	gio_hd = []string{
		"110100101100", "001101001011", "110011010010",
		"101100110100", "001011001101", "010010110011",
	}
	tiet_khi = []string{
		"Xuân Phân", "Thanh Minh", "Cốc Vũ", "Lập Hạ",
		"Tiểu Mãn", "Mang Chủng", "Hạ Chí", "Tiểu Thử",
		"Đại Thử", "Lập Thu", "Xử Thử", "Bạch Lộ",
		"Thu Phân", "Hàn Lộ", "Sương Giáng", "Lập Đông",
		"Tiểu Tuyết", "Đại Tuyết", "Đông Chí", "Tiểu Hàn",
		"Đại Hàn", "Lập Xuân", "Vũ Thủy", "Kinh Trập",
	}

	ngay_hoang_dao = map[int]map[string]string{
		1: {"Tý": "tốt", "Sửu": "tốt", "Tị": "tốt", "Mùi": "tốt", "Ngọ": "xấu", "Mão": "xấu", "Hợi": "xấu", "Dậu": "xấu"},
		2: {"Dần": "tốt", "Mão": "tốt", "Mùi": "tốt", "Dậu": "tốt", "Thân": "xấu", "Tị": "xấu", "Sửu": "xấu", "Hợi": "xấu"},
		3: {"Thìn": "tốt", "Tị": "tốt", "Dậu": "tốt", "Hợi": "tốt", "Tuất": "xấu", "Mùi": "xấu", "Sửu": "xấu"},
		4: {"Ngọ": "tốt", "Mùi": "tốt", "Sửu": "tốt", "Tý": "xấu", "Dậu": "xấu", "Tị": "xấu", "Mão": "xấu"},
		5: {"Thân": "tốt", "Dậu": "tốt", "Sửu": "tốt", "Mão": "tốt", "Dần": "xấu", "Hợi": "xấu", "Mùi": "xấu", "Tị": "xấu"},
		6: {"Tuất": "tốt", "Hợi": "tốt", "Mão": "tốt", "Tị": "tốt", "Thìn": "xấu", "Sửu": "xấu", "Dậu": "xấu", "Mùi": "xấu"},
	}

	thang_ngay_hoang_dao = map[int]int{
		1: 1, 7: 1, 2: 2, 8: 2, 3: 3, 9: 3, 4: 4, 10: 4, 5: 5, 11: 5, 6: 6, 12: 6,
	}
)

func getTietKhi(jd float64) string {
	return tiet_khi[getSolarTerm(jd+1, 7.0)]
}

func getYearCanChi(year int) string {
	return fmt.Sprintf("%s %s", can[(year+6)%10], chi[(year+8)%12])
}

func getMonthCanChi(mm int, yyyy int) string {
	return fmt.Sprintf("%s %s", can[(yyyy*12+mm+3)%10], chi[(mm+1)%12])
}

func getDayHoangHacDao(chi string, lunarDate LunarDate) string {
	monthIdx, ok := thang_ngay_hoang_dao[lunarDate.Month]
	if !ok {
		return ""
	}
	return ngay_hoang_dao[monthIdx][chi]
}

func getDayChi(jd int) string {
	return chi[(jd+1)%12]
}

func getDayCanChi(jd int, lunarDate LunarDate) string {
	can := can[(jd+9)%10]
	chi := getDayChi(jd)
	res := fmt.Sprintf("%s %s", can, chi)

	ngayHoangDao := getDayHoangHacDao(chi, lunarDate)
	if ngayHoangDao != "" {
		res += fmt.Sprintf(" (%s)", ngayHoangDao)
	}
	return res
}

func getCanHour0(jd int) string {
	return fmt.Sprintf("%s %s", can[((jd-1)*2)%10], chi[0])
}

func getGioHoangDao(jd int) string {
	chiOfDay := (jd + 1) % 12
	gioHD := gio_hd[chiOfDay%6]
	var parts []string
	count := 0

	for i := range 12 {
		if gioHD[i] == '1' {
			timeStr := fmt.Sprintf("%s (%d-%d)", chi[i], (i*2+23)%24, (i*2+1)%24)
			parts = append(parts, timeStr)
			count++
			if count >= 6 {
				break
			}
		}
	}
	return strings.Join(parts, ", ")
}

func getSolarTerm(dayNumber float64, timeZone float64) int {
	return int(math.Floor(sunLongitude(dayNumber-0.5-timeZone/24.0) / math.Pi * 12))
}

type CanChi struct {
	Year        string
	Month       string
	Day         string
	Hour        string
	TietKhi     string
	GioHoangDao string
}

func GetCanChi(vnDate VNDate) CanChi {
	jd := vnDate.Jd()
	return CanChi{
		Year:        getYearCanChi(vnDate.Year()),
		Month:       getMonthCanChi(int(vnDate.Month()), vnDate.Year()),
		Day:         getDayCanChi(jd, vnDate.LunarDate()),
		Hour:        getCanHour0(jd),
		TietKhi:     getTietKhi(float64(jd)),
		GioHoangDao: getGioHoangDao(jd),
	}
}
