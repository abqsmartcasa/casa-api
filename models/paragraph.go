package models

// Paragraph model for CASA paragraphs
type Paragraph struct {
	UUID            string `json:"-"`
	ID              int    `json:"id"`
	ParagraphNumber int    `json:"paragraphNumber"`
	ParagraphTitle  string `json:"paragraphTitle"`
	ParagraphText   string `json:"paragraphText"`
}

// AllParagraphs returns a slice with all paragraphs
func (db *DB) AllParagraphs(lang interface{}) ([]*Paragraph, error) {
	query := `SELECT
		paragraph."paragraph_number" AS "id",
		paragraph."paragraph_number" AS "paragraph_number",
		"trans_paragraph_description"."text" AS "paragraph_title",
		"trans_paragraph_text"."text" AS "paragraph_text"
	FROM
		paragraph
		LEFT JOIN "trans_paragraph_description"
		ON "trans_paragraph_description"."paragraph_uuid" = paragraph.uuid 
		LEFT JOIN "trans_paragraph_text"
		ON "trans_paragraph_text"."paragraph_uuid" = paragraph.uuid 
	WHERE
		"trans_paragraph_text"."lang_code" = $1`
	rows, err := db.Query(query, lang)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ps := make([]*Paragraph, 0)
	for rows.Next() {
		p := new(Paragraph)
		err := rows.Scan(&p.ID, &p.ParagraphNumber, &p.ParagraphTitle, &p.ParagraphText)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ps, nil
}

// GetParagraph Returns a single paragraph given a Paragraph.ID
func (db *DB) GetParagraph(lang interface{}, paragraph Paragraph) (*Paragraph, error) {
	p := new(Paragraph)
	query := `SELECT
			paragraph."paragraph_number" AS "id",
			paragraph."paragraph_number" AS "paragraph_number",
			"trans_paragraph_description"."text" AS "paragraph_title",
			"trans_paragraph_text"."text" AS "paragraph_text"
		FROM
			paragraph
			LEFT JOIN "trans_paragraph_description"
			ON "trans_paragraph_description"."paragraph_uuid" = paragraph.uuid 
			LEFT JOIN "trans_paragraph_text"
			ON "trans_paragraph_text"."paragraph_uuid" = paragraph.uuid 
		WHERE
			"trans_paragraph_text"."lang_code" = $1
			AND paragraph.paragraph_number = $2`
	row := db.QueryRow(query, lang, paragraph.ID)
	err := row.Scan(&p.ID, &p.ParagraphNumber, &p.ParagraphTitle, &p.ParagraphText)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// GetParagraphsBySpecificTag returns a slice with all paragraphs given a SpecificTag.ID
func (db *DB) GetParagraphsBySpecificTag(lang interface{}, specificTag SpecificTag) ([]*Paragraph, error) {
	query := `SELECT
			paragraph."paragraph_number" AS "id",
			paragraph."paragraph_number" AS "paragraph_number",
			"trans_paragraph_description"."text" AS "paragraph_title",
			"trans_paragraph_text"."text" AS "paragraph_text"
		FROM
			paragraph
			LEFT JOIN lkp_casa_specific
			ON lkp_casa_specific.id = paragraph.casa_specific_id
			LEFT JOIN lkp_casa_category
			ON lkp_casa_category.id = lkp_casa_specific.category_id
			LEFT JOIN "trans_paragraph_description"
			ON "trans_paragraph_description"."paragraph_uuid" = paragraph.uuid 
			LEFT JOIN "trans_paragraph_text"
			ON "trans_paragraph_text"."paragraph_uuid" = paragraph.uuid 
		WHERE
			"trans_paragraph_text"."lang_code" = $1
			AND lkp_casa_specific."id" = $2`
	rows, err := db.Query(query, lang, specificTag.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ps := make([]*Paragraph, 0)
	for rows.Next() {
		p := new(Paragraph)
		err := rows.Scan(&p.ID, &p.ParagraphNumber, &p.ParagraphTitle, &p.ParagraphText)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ps, nil
}

// GetParagraphsByCategoryTag returns a slice with all paragraphs given a CategoryTag.ID
func (db *DB) GetParagraphsByCategoryTag(lang interface{}, categoryTag CategoryTag) ([]*Paragraph, error) {
	query := `SELECT
			paragraph."paragraph_number" AS "id",
			paragraph."paragraph_number" AS "paragraph_number",
			"trans_paragraph_description"."text" AS "paragraph_title",
			"trans_paragraph_text"."text" AS "paragraph_text"
		FROM
			paragraph
			LEFT JOIN lkp_casa_specific
			ON lkp_casa_specific.id = paragraph.casa_specific_id
			LEFT JOIN lkp_casa_category
			ON lkp_casa_category.id = lkp_casa_specific.category_id
			LEFT JOIN "trans_paragraph_description"
			ON "trans_paragraph_description"."paragraph_uuid" = paragraph.uuid 
			LEFT JOIN "trans_paragraph_text"
			ON "trans_paragraph_text"."paragraph_uuid" = paragraph.uuid 
		WHERE
			"trans_paragraph_text"."lang_code" = $1
			AND lkp_casa_category."id" = $2`
	rows, err := db.Query(query, lang, categoryTag.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ps := make([]*Paragraph, 0)
	for rows.Next() {
		p := new(Paragraph)
		err := rows.Scan(&p.ID, &p.ParagraphNumber, &p.ParagraphTitle, &p.ParagraphText)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ps, nil
}

/*
func GetParagraphsBySOP(lang interface{}, sop SOP) ([]*Paragraph, error) {

}
*/
