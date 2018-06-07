package utils

import (
	_ "github.com/lib/pq"
	"fmt"
	"bytes"
	"github.com/go-pg/pg"
	"os"
	"unicode"
	"bufio"
	"strings"
	"strconv"
	"time"
	"regexp"

	"io/ioutil"
	"github.com/tkanos/gonfig"
	"io"
	//"gopkg.in/cheggaaa/pb.v1"
	//"log"
	"reflect"
)

//db struct
type Houses struct {
	tableName struct{} `sql:"vmetre.houses"`
	//Unic_id							int
	Id                              int     //ID дома на Портале
	Region_id                       string  //`Субъект РФ (код ФИАС)`
	Area_id                         string  //`Район (код ФИАС)`
	City_id                         string  //`Населенный пункт (код ФИАС)`
	Street_id                       string  //`Улица (код ФИАС)`
	Shortname_region                string  //`Тип Субъекта РФ`
	Formalname_region               string  //`Субъект РФ (наименование)`
	Shortname_area                  string  //`Тип района`
	Formalname_area                 string  //`Район (наименование)`
	Shortname_city                  string  //`Тип населенного пункта`
	Formalname_city                 string  //`Населенный пункт (наименование)`
	Shortname_street                string  //`Тип улицы`
	Formalname_street               string  //`Улица (наименование)`
	House_number                    string  //`Номер дома`
	Building                        string  //`Корпус`
	Block                           string  //`Строение`
	Letter                          string  //`Литера`
	Address                         string  //`Адрес дома`
	Houseguid                       string  //`Глобальный уникальный идентификатор дома`
	Management_organization_id      string  //`ID Управляющей организации на Портале`
	Built_year                      string  //`Год постройки`
	Exploitation_start_year         string  //`Год ввода в эксплуатацию`
	Project_type                    string  //`Серия, тип постройки здания`
	House_type                      string  //`Тип дома`
	Is_alarm                        string  //`Факт признания дома аварийным`
	Method_of_forming_overhaul_fund string  //`Способ формирования фонда капитального ремонта`
	Floor_count_max                 int     //`Наибольшее количество этажей, ед.`
	Floor_count_min                 int     //`Наименьшее количество этажей, ед.`
	Entrance_count                  int     //`Количество подъездов, ед.`
	Elevators_count                 int     //`Количество лифтов, ед.`
	Energy_efficiency               string  //`Класс энергетической эффективности`
	Quarters_count                  int     //`Количество помещений, всего, ед.`
	Living_quarters_count           int     //`Количество жилых помещений, ед.`
	Unliving_quarters_count         int     //`Количество нежилых помещений, ед.`
	Area_total                      float32 //`Общая площадь дома, всего, кв.м`
	Area_residential                float32 //`Общая площадь жилых помещений, кв.м`
	Area_non_residential            float32 //`Общая площадь нежилых помещений, кв.м`
	Area_common_property            float32 //`Общая площадь помещений, входящих в состав общего имущества, кв.м`
	Area_land                       float32 //`Площадь земельного участка, входящего в состав общего имущества в многоквартирном доме, кв.м`
	Parking_square                  float32 //`Площадь парковки в границах земельного участка, кв.м`
	Playground                      string  //`Элементы благоустройства (детская площадка)`
	Sportsground                    string  //`Элементы благоустройства (спортивная площадка)`
	Other_beautification            string  //`Элементы благоустройства (другое)`
	Foundation_type                 string  //`Тип фундамента`
	Floor_type                      string  //`Тип перекрытий`
	Wall_material                   string  //`Материал несущих стен`
	Basement_area                   float32 //`Площадь подвала по полу, кв.м`
	Chute_type                      string  //`Тип мусоропровода`
	Chute_count                     int     //`Количество мусоропроводов, ед.`
	Electrical_type                 string  //`Тип системы электроснабжения`
	Electrical_entries_count        int     //`Количество вводов в МКД, ед.`
	Heating_type                    string  //`Тип системы теплоснабжения`
	Hot_water_type                  string  //`Тип системы горячего водоснабжения`
	Cold_water_type                 string  //	`Тип системы холодного водоснабжения`
	Sewerage_type                   string  //`Тип системы водоотведения`
	Sewerage_cesspools_volume       float32 //`Объем выгребных ям, куб.м`
	Gas_type                        string  //	`Тип системы газоснабжения`
	Ventilation_type                string  //	`Тип системы вентиляции`
	Firefighting_type               string  //`Тип системы пожаротушения`
	Drainage_type                   string  //`Тип системы водостоков`
	Uid                             int
}
type UprOrg struct {
	tableName struct{} `sql:"vmetre.uprorg"`
	//Unic_id							int
	Id             int     //ID Управляющей организации на Портале
	Subject_rf     string  //Субъект РФ
	Name_full      string  //Фирменное наименование юридического лица (согласно уставу организации)
	Name_short     string  //Сокращенное наименование
	Name_employee  string  //ФИО руководителя
	Inn            int     //Идентификационный номер налогоплательщика (ИНН)
	Orn            int     //ОГРН
	Legal_address  string  //Место государственной регистрации юридического лица (адрес юридического лица)
	Actual_address string  //Адрес фактического местонахождения органов управления
	Phone          string  //Контактные телефоны
	Email          string  //Адрес электронной почты
	Site           string  //Официальный сайт в сети Интернет
	Count_mkd      int     //Количество домов, находящихся в управлении, ед.
	Area_total     float32 //Площадь домов, находящихся в управлении, кв. м
	W_summ         float32 //Присвоенный рейтинг на дату выгрузки
	Uid            int
}
type UpdatedFiles struct {
	tableName struct{} `sql:"vmetre.filecontrol"`
	Uid       int      `sql:",pk"`
	FileName  string
	StartTime time.Time
	EndTime   pg.NullTime //*time.Time
	RowsCount int
}
type Configuration struct {
	User     string
	Password string
	Database string
	Addr     string
}

var configuration Configuration

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
func ToDB() {
	err := gonfig.GetConf("utils/config.json", &configuration)
	if err != nil {
		fmt.Println(err)
	}
	db := pg.Connect(&pg.Options{
		User:     configuration.User,
		Password: configuration.Password,
		Database: configuration.Database,
		Addr:     configuration.Addr,
	})
	defer db.Close()
	initDatabase(&db)
	files, _ := ioutil.ReadDir("output/")

	var dataStruct interface{}
	for _, f := range files {
		uprorgs := regexp.MustCompile("(reestruo*)")
		houses := regexp.MustCompile("(reestrmkd*)")

		fmt.Println("Goes through: " + f.Name())
		file := f.Name()
		if uprorgs.MatchString(file) {
			dataStruct = &UprOrg{}
		} else if houses.MatchString(file) {
			dataStruct = &Houses{}
		}
		test := UpdatedFiles{
			FileName:  file,
			StartTime: time.Now(),
		}
		var file1 UpdatedFiles
		count, _ := db.Model(&file1).Where("file_name = ?", file).SelectAndCount()
		if count != 0 {

			if file1.EndTime.IsZero() {

				writeToDb(file, db, dataStruct, file1, file1.RowsCount)
				err := os.Rename("output/"+file, "finished/"+file)
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				fmt.Println(file + " is the same")
				err := os.Rename("output/"+file, "finished/"+file)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		} else {
			db.Insert(&test)
			writeToDb(file, db, dataStruct, test, 0)
			err := os.Rename("output/"+file, "finished/"+file)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

//trim first byte of csv file for correct reading
func trapBOM(fileBytes []byte) []byte {
	trimmedBytes := bytes.Trim(fileBytes, "\xef\xbb\xbf")
	return trimmedBytes
}

//up a first letter of key in row what allows to set value for struct
func fieldNameColumn(typ string) (string) {
	a := []rune(typ)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}

//put values from file to db
//first arg is name of file to extract
//second is db to write into
//last arg is type of struct: 1 is for Houses struct, 2 is for UprOrg struct
func writeToDb(nameFile string, db *pg.DB, dataType interface{}, files UpdatedFiles, row int) {
	//rowsCount := 0
	file, err := os.Open("output/" + nameFile)
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()
	//lines,err:= lineCounter(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)
	scanner1 := bufio.NewReader(file)
	line, _, _ := scanner1.ReadLine()
	firstRow:= strings.Split(string(line), ";")
	firstRow[0] = string(trapBOM([]byte(firstRow[0])))

	var state = 1
	var result []string
	//var begin int
	//var end int
	var r []rune

	for {
		tok, _, err := scanner1.ReadRune();
		//fmt.Println(string(tok), state)
		if err == io.EOF  {
			break
		}
		switch tok {
		case '"':
			if state == 2 {
				state = 4
			} else if state == 4 {
				r = append(r, tok)
				state = 5
			} else if state == 5 {
				r = append(r, tok)
				state = 4
			}
		case '\n':
			if state == 2 ||state==5{
				state = 1
				if r[len(r)-1]=='"'{
					r=r[:len(r)-1]
				}
				//r = append(r, tok)
				result = append(result, string(r))
				//fmt.Println(" to chan ", result)
				Save(result, firstRow,dataType,*db,files)
				result=nil
			}
		case '\r':


		case ';':
			if state == 2 {
				result = append(result, string(r))
				r=nil
			} else if state == 5 {
				if r[len(r)-1]=='"'{
					r=r[:len(r)-1]
				}
				result = append(result, string(r))
				r=nil
				state = 2
			}
		default:
			if state == 1 {
				r = nil
				r = append(r, tok)
				state = 2
			} else if state == 5 {
				r = append(r, tok)
				state = 2
			} else if state == 4 {
				r = append(r, tok)
			} else if state == 2 {
				r = append(r, tok)
			}

		}
	}
	if r[len(r)-1]=='"'{
		r=r[:len(r)-1]
	}
	//result = append(result, string(r))
	//fmt.Println(" to chan ")
	//Save(result, firstRow,dataType,*db,files)
	//result=nil
}

func Save(x []string, columns []string, dataType interface{}, db pg.DB, files UpdatedFiles){
	//fmt.Println(x)
	if len(x)==0{
		return
	}
	columnLen := len(columns)
	//fmt.Println(" from chan ",x)
	for i := 0; i < columnLen; i++ {
		fieldName := fieldNameColumn(columns[i])
		SetValue(fieldName, x[i], &dataType)
	}
	SetValue(fieldNameColumn("Uid"), strconv.Itoa(files.Uid), &dataType)
	err := db.Insert(dataType)
	if err != nil {
		panic(err)
	}
	files.EndTime = pg.NullTime{time.Now()}
	_, error1 := db.Model(&files).Column("end_time").Update(&files)
	if error1 != nil {
		panic(error1)
	}

}


//make a slice of value for struct insert
func Cut(str []rune, b int, e int) string {
	str = str[b:e]
	value := string(str)
	return strings.TrimSpace(value)
}

//set value using reflection
func SetValue(fieldName string, value string, obj *interface{}) {
	f := reflect.ValueOf(*obj).Elem().FieldByName(fieldName)

	switch f.Kind() {
	case reflect.Int:
		val, err := strconv.ParseInt(value, 0, 64)
		if err == nil {
			f.SetInt(val)

		}
	case reflect.String:

		f.SetString(value)

	case reflect.Float32:
		v:=strings.Trim(strings.Replace(value, ",", ".", -1)," ")
		val, err := strconv.ParseFloat(v, 32)
		if err == nil {
			f.SetFloat(val)
		}else		{
			if len(v)>0 {
				panic(err)
			}
		}

	}
}

//sql command for creating database
func initDatabase(pdb **pg.DB) {
	db := *pdb
	initHouseDb := []string{
		"CREATE SCHEMA IF NOT EXISTS ******;",
		`CREATE TABLE IF NOT EXISTS ******.houses
(
    id integer NOT NULL,
    region_id varchar,
    area_id varchar,
    city_id varchar,
    street_id varchar,
    shortname_region varchar,
    formalname_region varchar,
    shortname_area varchar,
    formalname_area varchar,
    shortname_city varchar,
    formalname_city varchar,
    shortname_street varchar,
    formalname_street varchar,
    house_number varchar,
    building varchar,
    block varchar,
    letter varchar,
    address varchar,
    houseguid varchar,
    management_organization_id varchar,
    built_year integer,
    exploitation_start_year varchar,
    project_type varchar,
    house_type varchar,
    is_alarm varchar,
    method_of_forming_overhaul_fund varchar,
    floor_count_max integer,
    floor_count_min integer,
    entrance_count integer,
    elevators_count integer,
    energy_efficiency varchar,
    quarters_count integer,
    living_quarters_count integer,
    unliving_quarters_count integer,
    area_total double precision,
    area_residential double precision,
    area_non_residential double precision,
    area_common_property double precision,
    area_land double precision,
    parking_square double precision,
    playground varchar,
    sportsground varchar,
    other_beautification varchar,
    foundation_type varchar,
    floor_type varchar,
    wall_material varchar,
    basement_area double precision,
    chute_type varchar,
    chute_count integer,
    electrical_type varchar,
    electrical_entries_count integer,
    heating_type varchar,
    hot_water_type varchar,
    cold_water_type varchar,
    sewerage_type varchar,
    sewerage_cesspools_volume double precision,
    gas_type varchar,
    ventilation_type varchar,
    firefighting_type varchar,
    drainage_type varchar,
	uid int REFERENCES vmetre.filecontrol);`}

	initUprOgsDb := []string{
		"CREATE SCHEMA IF NOT EXISTS ******;",
		`CREATE TABLE IF NOT EXISTS ******.uprorg
		(
			id             INTEGER NOT NULL,    
			subject_rf     VARCHAR,  
			name_full      VARCHAR,  
			name_short     VARCHAR, 
			name_employee  VARCHAR, 
			inn            INTEGER,    
			ogrn           INTEGER,    
			legal_address  VARCHAR,  
			actual_address VARCHAR,  
			phone          VARCHAR, 
			email          VARCHAR,  
			site           VARCHAR,  
			count_mkd      VARCHAR,
			area_total     double precision,
			rating_uo      double precision,
			uid 		   int REFERENCES vmetre.filecontrol
			);`}

	initFileControl := []string{
		"CREATE SCHEMA IF NOT EXISTS vmetre;",
		`CREATE TABLE IF NOT EXISTS vmetre.filecontrol
		(
			uid SERIAL PRIMARY KEY,
			file_name varchar,
			start_time time,
			end_time time,
			rows_count integer);`}

	for _, str := range initFileControl {
		_, err := db.Exec(str)
		if err != nil {
			fmt.Printf("Database init error -->%v\n", err)
			fmt.Printf("Query error")
		}
	}

	for _, str := range initHouseDb {
		_, err := db.Exec(str)
		if err != nil {
			fmt.Printf("Database init error -->%v\n", err)
			fmt.Printf("Query error")
		}
	}

	for _, str := range initUprOgsDb {
		_, err := db.Exec(str)
		if err != nil {
			fmt.Printf("Database init error -->%v\n", err)
			fmt.Printf("Query error")
		}
	}

}