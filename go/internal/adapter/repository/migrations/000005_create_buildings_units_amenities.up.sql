-- ENUMs
CREATE TYPE building_status AS ENUM ('draft', 'published', 'archived');
CREATE TYPE unit_status AS ENUM ('draft', 'available', 'reserved', 'rented', 'archived');
CREATE TYPE amenity_category AS ENUM (
    'interior', 'kitchen', 'bathroom', 'outdoor',
    'building', 'technology', 'parking', 'accessibility', 'other'
);
CREATE TYPE media_entity_type AS ENUM ('building', 'unit');

-- Buildings (Gebäude/Objekte)
CREATE TABLE IF NOT EXISTS buildings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    street VARCHAR(255) NOT NULL,
    house_number VARCHAR(50) NOT NULL,
    postal_code VARCHAR(20) NOT NULL,
    city VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL DEFAULT 'DE',
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    year_built INT,
    total_floors INT,
    total_units INT DEFAULT 0,
    description TEXT,
    status building_status  NOT NULL DEFAULT 'draft',
    published BOOLEAN NOT NULL DEFAULT FALSE,
    published_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT unique_tenant_building_slug UNIQUE (tenant_id, slug)
);

CREATE INDEX idx_buildings_tenant_id ON buildings(tenant_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_buildings_slug ON buildings(slug) WHERE deleted_at IS NULL;
CREATE INDEX idx_buildings_city ON buildings(city) WHERE deleted_at IS NULL;
CREATE INDEX idx_buildings_status ON buildings(status) WHERE deleted_at IS NULL;

ALTER TABLE buildings ENABLE ROW LEVEL SECURITY;

CREATE POLICY tenant_buildings_isolation_policy ON buildings
    USING (tenant_id = current_setting('app.current_tenant_id')::UUID);


-- Units (Wohnungen/Einheiten)
CREATE TABLE IF NOT EXISTS units (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    building_id UUID NOT NULL REFERENCES buildings(id) ON DELETE CASCADE,
    unit_number VARCHAR(50) NOT NULL,
    floor INT,
    rooms DECIMAL(3, 1) NOT NULL,
    bedrooms INT NOT NULL DEFAULT 0,
    bathrooms DECIMAL(3, 1) NOT NULL DEFAULT 1,
    living_area_sqm DECIMAL(8, 2),
    total_area_sqm DECIMAL(8, 2),
    balcony_area_sqm DECIMAL(8, 2),
    ceiling_height_m DECIMAL(4, 2),
    cold_rent DECIMAL(10, 2),
    warm_rent DECIMAL(10, 2),
    deposit DECIMAL(10, 2),
    available_from DATE,
    status unit_status NOT NULL DEFAULT 'draft',
    published BOOLEAN NOT NULL DEFAULT FALSE,
    published_at TIMESTAMP WITH TIME ZONE,
    title VARCHAR(500),
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT unique_building_unit_number UNIQUE (building_id, unit_number)
);

CREATE INDEX idx_units_tenant_id ON units(tenant_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_units_building_id ON units(building_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_units_status ON units(status) WHERE deleted_at IS NULL;
CREATE INDEX idx_units_available_from ON units(available_from) WHERE deleted_at IS NULL;

ALTER TABLE units ENABLE ROW LEVEL SECURITY;

CREATE POLICY tenant_units_isolation_policy ON units
    USING (tenant_id = current_setting('app.current_tenant_id')::UUID);


-- Amenities (Ausstattungsmerkmale)
CREATE TABLE IF NOT EXISTS amenities (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID REFERENCES tenants(id) ON DELETE CASCADE,
    code VARCHAR(100) NOT NULL,
    name VARCHAR(255) NOT NULL,
    category amenity_category NOT NULL,
    icon VARCHAR(100),
    is_global BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_tenant_amenity_code UNIQUE (tenant_id, code)
);

CREATE INDEX idx_amenities_tenant_id ON amenities(tenant_id);
CREATE INDEX idx_amenities_category ON amenities(category);
CREATE INDEX idx_amenities_is_global ON amenities(is_global);


-- Unit-Amenities (n:m Verknüpfung)
CREATE TABLE IF NOT EXISTS unit_amenities (
    unit_id UUID NOT NULL REFERENCES units(id) ON DELETE CASCADE,
    amenity_id UUID NOT NULL REFERENCES amenities(id) ON DELETE CASCADE,
    additional_info VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY (unit_id, amenity_id)
);

CREATE INDEX idx_unit_amenities_unit_id ON unit_amenities(unit_id);
CREATE INDEX idx_unit_amenities_amenity_id ON unit_amenities(amenity_id);


-- Building-Amenities (n:m Verknüpfung)
CREATE TABLE IF NOT EXISTS building_amenities (
    building_id UUID NOT NULL REFERENCES buildings(id) ON DELETE CASCADE,
    amenity_id UUID NOT NULL REFERENCES amenities(id) ON DELETE CASCADE,
    additional_info VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY (building_id, amenity_id)
);

CREATE INDEX idx_building_amenities_building_id ON building_amenities(building_id);
CREATE INDEX idx_building_amenities_amenity_id ON building_amenities(amenity_id);


-- Media (Bilder/Dokumente für Buildings und Units)
CREATE TABLE IF NOT EXISTS media (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    entity_type media_entity_type NOT NULL,
    entity_id UUID NOT NULL,
    file_name VARCHAR(500) NOT NULL,
    file_path VARCHAR(1000) NOT NULL,
    file_size BIGINT NOT NULL,
    mime_type VARCHAR(100) NOT NULL,
    alt_text VARCHAR(500),
    sort_order INT NOT NULL DEFAULT 0,
    is_primary BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT valid_entity_type CHECK (entity_type IN ('building', 'unit'))
);

CREATE INDEX idx_media_entity ON media(entity_type, entity_id);
CREATE INDEX idx_media_tenant_id ON media(tenant_id);