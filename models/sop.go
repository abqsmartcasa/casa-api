package models

// SOP model for Independent monitoring reports
type SOP struct {
	UUID          string `json:"-"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	Title         string `json:"title"`
	Current       bool   `json:"current"`
	EffectiveDate string `json:"effective_date"`
	ReviewDate    string `json:"review_date"`
	ReplacesDate  string `json:"replaces_date"`
}

// AllSOPS returns a slice of all SOPS
func (db *DB) AllSOPs(lang interface{}) ([]*SOP, error) {
	query := ``
	rows, err := db.Query(query, lang)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sops := make([]*SOP, 0)
	for rows.Next() {
		s := new(SOP)
		err := rows.Scan(&s.ID, &s.Name, &s.Title, &s.EffectiveDate, &s.ReviewDate, &s.ReplacesDate)
		if err != nil {
			return nil, err
		}
		sops = append(sops, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return sops, nil
}

// GetSOP returns a single SOP given a SOP.ID
func (db *DB) GetSOP(lang interface{}, sop SOP) (*SOP, error) {
	s := new(SOP)
	query := ``
	row := db.QueryRow(query, lang, sop.ID)
	err := row.Scan(&s.ID, &s.Name, &s.Title, &s.EffectiveDate, &s.ReviewDate, &s.ReplacesDate)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetSOPSByParagraph returns SOPs given a paragraph.ID
func (db *DB) GetSOPsByParagraph(lang interface{}, paragraph Paragraph) ([]*SOP, error) {
	query := ``
	rows, err := db.Query(query, lang, paragraph)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sops := make([]*SOP, 0)
	for rows.Next() {
		s := new(SOP)
		err := rows.Scan(&s.ID, &s.Name, &s.Title, &s.EffectiveDate, &s.ReviewDate, &s.ReplacesDate)
		if err != nil {
			return nil, err
		}
		sops = append(sops, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return sops, nil
}

// GetSOPsBySpecficTag returns SOPs given a specificTag.ID
func (db *DB) GetSOPsBySpecificTag(lang interface{}, specificTag SpecificTag) ([]*SOP, error) {
	query := ``
	rows, err := db.Query(query, lang, specificTag.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sops := make([]*SOP, 0)
	for rows.Next() {
		s := new(SOP)
		err := rows.Scan(&s.ID, &s.Name, &s.Title, &s.EffectiveDate, &s.ReviewDate, &s.ReplacesDate)
		if err != nil {
			return nil, err
		}
		sops = append(sops, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return sops, nil
}

// GetSOPsByCategoryTag returns SOPs given a categoryTag.ID
func (db *DB) GetSOPsByCategoryTag(lang interface{}, categoryTag CategoryTag) ([]*SOP, error) {
	query := ``
	rows, err := db.Query(query, lang, categoryTag.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sops := make([]*SOP, 0)
	for rows.Next() {
		s := new(SOP)
		err := rows.Scan(&s.ID, &s.Name, &s.Title, &s.EffectiveDate, &s.ReviewDate, &s.ReplacesDate)
		if err != nil {
			return nil, err
		}
		sops = append(sops, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return sops, nil
}
