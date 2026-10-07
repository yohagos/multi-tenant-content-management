-- go/internal/adapter/repository/migrations/006_seed_amenities.up.sql
INSERT INTO amenities (code, name, category, icon, is_global) VALUES
-- Interior
('parquet', 'Parkettboden', 'interior', 'floor', TRUE),
('laminate', 'Laminatboden', 'interior', 'floor', TRUE),
('tiles', 'Fliesenboden', 'interior', 'floor', TRUE),
('underfloor_heating', 'Fußbodenheizung', 'interior', 'heat', TRUE),
('fireplace', 'Kamin', 'interior', 'fire', TRUE),
('furnished', 'Möbliert', 'interior', 'sofa', TRUE),
('built_in_wardrobe', 'Einbauschränke', 'interior', 'wardrobe', TRUE),

-- Kitchen
('fitted_kitchen', 'Einbauküche', 'kitchen', 'kitchen', TRUE),
('dishwasher', 'Spülmaschine', 'kitchen', 'dishwasher', TRUE),
('fridge', 'Kühlschrank', 'kitchen', 'fridge', TRUE),
('oven', 'Backofen', 'kitchen', 'oven', TRUE),
('microwave', 'Mikrowelle', 'kitchen', 'microwave', TRUE),
('induction_stove', 'Induktionsherd', 'kitchen', 'stove', TRUE),

-- Bathroom
('bathtub', 'Badewanne', 'bathroom', 'bathtub', TRUE),
('shower', 'Dusche', 'bathroom', 'shower', TRUE),
('guest_toilet', 'Gäste-WC', 'bathroom', 'toilet', TRUE),
('washing_machine', 'Waschmaschine', 'bathroom', 'washing-machine', TRUE),
('dryer', 'Trockner', 'bathroom', 'dryer', TRUE),
('towel_radiator', 'Handtuchheizkörper', 'bathroom', 'radiator', TRUE),

-- Outdoor
('balcony', 'Balkon', 'outdoor', 'balcony', TRUE),
('terrace', 'Terrasse', 'outdoor', 'terrace', TRUE),
('garden', 'Garten', 'outdoor', 'garden', TRUE),
('loggia', 'Loggia', 'outdoor', 'balcony', TRUE),
('roof_terrace', 'Dachterrasse', 'outdoor', 'terrace', TRUE),

-- Building
('elevator', 'Aufzug', 'building', 'elevator', TRUE),
('basement', 'Keller', 'building', 'basement', TRUE),
('attic', 'Dachboden', 'building', 'attic', TRUE),
('concierge', 'Concierge', 'building', 'concierge', TRUE),
('bicycle_storage', 'Fahrradkeller', 'building', 'bicycle', TRUE),
('laundry_room', 'Waschküche', 'building', 'laundry', TRUE),
('common_room', 'Gemeinschaftsraum', 'building', 'community', TRUE),

-- Technology
('fiber_internet', 'Glasfaseranschluss', 'technology', 'internet', TRUE),
('cable_tv', 'Kabelanschluss', 'technology', 'tv', TRUE),
('smart_home', 'Smart Home', 'technology', 'smart-home', TRUE),
('air_conditioning', 'Klimaanlage', 'technology', 'ac', TRUE),
('solar_panels', 'Solaranlage', 'technology', 'solar', TRUE),
('ev_charging', 'E-Ladestation', 'technology', 'ev-charging', TRUE),

-- Parking
('parking_space', 'Stellplatz', 'parking', 'parking', TRUE),
('garage', 'Garage', 'parking', 'garage', TRUE),
('underground_parking', 'Tiefgarage', 'parking', 'underground', TRUE),

-- Accessibility
('wheelchair_accessible', 'Barrierefrei', 'accessibility', 'wheelchair', TRUE),
('step_free_access', 'Stufenloser Zugang', 'accessibility', 'accessible', TRUE),
('wide_doors', 'Verbreiterte Türen', 'accessibility', 'door', TRUE),

-- Other
('pet_friendly', 'Haustiere erlaubt', 'other', 'pet', TRUE),
('smoking_allowed', 'Rauchen erlaubt', 'other', 'smoking', TRUE);