package models

type Base struct {

}

func (b *Base) TableEngine() string {
	return "INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci"
}
