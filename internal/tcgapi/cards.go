package tcgapi

func (c *Client) GetCard(id string) (*Card, error) {
	var card Card

	if err := c.get("cards/"+id, &card); err != nil {
		return nil, err
	}
	return &card, nil
}
