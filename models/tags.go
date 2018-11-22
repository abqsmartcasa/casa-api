package models

import "strings"

// SpecificTag model for specific tags
type SpecificTag struct {
	ID         int    `json:"id"`
	Value      string `json:"value"`
	CategoryID int    `json:"category_id,omitempty"`
}

// CategoryTag model for category tags
type CategoryTag struct {
	ID           int            `json:"id"`
	Value        string         `json:"value"`
	SpecificTags []*SpecificTag `json:"specific_tags,omitempty"`
}

var categoryTagQuery = `SELECT
			"lkp_casa_category"."id",
			"trans_casa_category"."text" AS "value"
		FROM
			"trans_casa_category"
			JOIN "lkp_casa_category"
			ON "trans_casa_category"."casa_category_id" = "lkp_casa_category"."id"
		WHERE
			"trans_casa_category"."lang_code" = $1`

var specificTagQuery = `SELECT
			"lkp_casa_specific"."id",
			"trans_casa_specific"."text" AS "value",
			"lkp_casa_specific"."category_id"
		FROM
			"trans_casa_specific"
			JOIN "lkp_casa_specific"
			ON "trans_casa_specific"."casa_specific_id" = "lkp_casa_specific"."id"
		WHERE
			"trans_casa_specific"."lang_code" = $1`

// AllCategoryTags returns a slice with all Category Tags
func (db *DB) AllCategoryTags(lang interface{}) ([]*CategoryTag, error) {
	rows, err := db.Query(categoryTagQuery, lang)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cts := make([]*CategoryTag, 0)
	for rows.Next() {
		ct := new(CategoryTag)
		err := rows.Scan(&ct.ID, &ct.Value)
		if err != nil {
			return nil, err
		}
		cts = append(cts, ct)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cts, nil
}

// GetCategoryTag returns a Category Tag given a CategoryTag.ID
func (db *DB) GetCategoryTag(lang interface{}, categoryTag CategoryTag) (*CategoryTag, error) {
	ct := new(CategoryTag)
	var ctQueryBuilder strings.Builder
	ctQueryBuilder.WriteString(categoryTagQuery)
	ctQueryBuilder.WriteString(" AND lkp_casa_category.id = $2")
	query := ctQueryBuilder.String()
	row := db.QueryRow(query, lang, categoryTag.ID)
	err := row.Scan(&ct.ID, &ct.Value)
	if err != nil {
		return nil, err
	}
	return ct, nil
}

// AllSpecificTags returns a slice with all Specific Tags
func (db *DB) AllSpecificTags(lang interface{}) ([]*SpecificTag, error) {
	rows, err := db.Query(specificTagQuery, lang)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sts := make([]*SpecificTag, 0)
	for rows.Next() {
		st := new(SpecificTag)
		err := rows.Scan(&st.ID, &st.Value)
		if err != nil {
			return nil, err
		}
		sts = append(sts, st)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return sts, nil
}

// GetSpecificTag returns a Category Tag given a SpecificTag.ID
func (db *DB) GetSpecificTag(lang interface{}, specificTag SpecificTag) (*SpecificTag, error) {
	st := new(SpecificTag)
	var stQueryBuilder strings.Builder
	stQueryBuilder.WriteString(specificTagQuery)
	stQueryBuilder.WriteString(" AND lkp_casa_specific.id = $2")
	query := stQueryBuilder.String()
	row := db.QueryRow(query, lang, specificTag.ID)
	err := row.Scan(&st.ID, &st.Value)
	if err != nil {
		return nil, err
	}
	return st, nil
}
