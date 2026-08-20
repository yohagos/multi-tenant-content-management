CREATE TABLE IF NOT EXISTS contents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    title VARCHAR(500) NOT NULL,
    slug VARCHAR(500) NOT NULL,
    body TEXT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'draft',
    published BOOLEAN NOT NULL DEFAULT FALSE,
    published_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_tenant FOREIGN KEY (tenant_id) REFERENCES tenants(id) ON DELETE CASCADE,
    CONSTRAINT unique_tenant_slug UNIQUE (tenant_id, slug)
);

CREATE INDEX idx_contents_tenant_id ON contents(tenant_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_contents_slug ON contents(slug) WHERE deleted_at IS NULL;
CREATE INDEX idx_contents_status ON contents(status) WHERE deleted_at IS NULL;
CREATE INDEX idx_contents_published ON contents(published) WHERE deleted_at IS NULL;
CREATE INDEX idx_contents_deleted_at ON contents(deleted_at);
CREATE INDEX idx_contents_tenant_status ON contents(tenant_id, status) WHERE deleted_at IS NULL;
CREATE INDEX idx_contents_tenant_published ON contents(tenant_id, published) WHERE deleted_at IS NULL;

ALTER TABLE contents ENABLE ROW LEVEL SECURITY;

CREATE POLICY tenant_content_isolation_policy ON contents
    USING (tenant_id = current_setting('app.current_tenant_id')::UUID);