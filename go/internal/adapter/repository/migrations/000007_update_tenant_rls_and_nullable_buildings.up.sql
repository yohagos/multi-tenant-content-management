BEGIN;

-- =========================================================
-- 1. Buildings and units may exist without a tenant
-- =========================================================

ALTER TABLE buildings
    ALTER COLUMN tenant_id DROP NOT NULL;

ALTER TABLE units
    ALTER COLUMN tenant_id DROP NOT NULL;

-- =========================================================
-- 2. Users: enable RLS
-- =========================================================

ALTER TABLE users ENABLE ROW LEVEL SECURITY;

-- =========================================================
-- 3. Replace tenant policies
-- =========================================================

DROP POLICY IF EXISTS tenant_isolation_policy ON tenants;
DROP POLICY IF EXISTS tenant_content_isolation_policy ON contents;
DROP POLICY IF EXISTS tenant_users_isolation_policy ON users;
DROP POLICY IF EXISTS tenant_buildings_isolation_policy ON buildings;
DROP POLICY IF EXISTS tenant_units_isolation_policy ON units;

-- =========================================================
-- 4. Tenant context and role-based access
--
-- Expected transaction-local settings:
-- app.current_tenant_id = UUID of the current tenant
-- app.current_role      = admin | tenant | editor | public
-- =========================================================

-- Tenants:
-- admin/editor: all tenants
-- tenant: own tenant
-- public: no access

CREATE POLICY tenants_select_policy ON tenants
    FOR SELECT
    USING (
        current_setting('app.current_role', true) IN ('admin', 'editor')
        OR (
            current_setting('app.current_role', true) = 'tenant'
            AND id = NULLIF(
                current_setting('app.current_tenant_id', true), ''
            )::UUID
        )
    );

CREATE POLICY tenants_modify_policy ON tenants
    FOR ALL
    USING (
        current_setting('app.current_role', true) IN ('admin', 'editor')
    )
    WITH CHECK (
        current_setting('app.current_role', true) IN ('admin', 'editor')
    );

-- Contents:
-- admin/editor: all contents
-- tenant: own tenant's contents
-- public: published contents

CREATE POLICY contents_select_policy ON contents
    FOR SELECT
    USING (
        current_setting('app.current_role', true) IN ('admin', 'editor')
        OR (
            current_setting('app.current_role', true) = 'tenant'
            AND tenant_id = NULLIF(
                current_setting('app.current_tenant_id', true), ''
            )::UUID
        )
        OR (
            current_setting('app.current_role', true) = 'public'
            AND published = TRUE
            AND deleted_at IS NULL
        )
    );

CREATE POLICY contents_modify_policy ON contents
    FOR ALL
    USING (
        current_setting('app.current_role', true) IN ('admin', 'editor')
    )
    WITH CHECK (
        current_setting('app.current_role', true) IN ('admin', 'editor')
    );

-- Users:
-- admin/editor: all users
-- tenant: users belonging to own tenant
-- public: no access

CREATE POLICY users_select_policy ON users
    FOR SELECT
    USING (
        current_setting('app.current_role', true) IN ('admin', 'editor')
        OR (
            current_setting('app.current_role', true) = 'tenant'
            AND tenant_id = NULLIF(
                current_setting('app.current_tenant_id', true), ''
            )::UUID
        )
    );

CREATE POLICY users_modify_policy ON users
    FOR ALL
    USING (
        current_setting('app.current_role', true) IN ('admin', 'editor')
    )
    WITH CHECK (
        current_setting('app.current_role', true) IN ('admin', 'editor')
    );

-- Buildings:
-- admin/editor: all buildings, including unassigned buildings
-- tenant: buildings assigned to current tenant
-- public: published buildings only

CREATE POLICY buildings_select_policy ON buildings
    FOR SELECT
    USING (
        current_setting('app.current_role', true) IN ('admin', 'editor')
        OR (
            current_setting('app.current_role', true) = 'tenant'
            AND tenant_id = NULLIF(
                current_setting('app.current_tenant_id', true), ''
            )::UUID
        )
        OR (
            current_setting('app.current_role', true) = 'public'
            AND published = TRUE
            AND deleted_at IS NULL
        )
    );

CREATE POLICY buildings_modify_policy ON buildings
    FOR ALL
    USING (
        current_setting('app.current_role', true) IN ('admin', 'editor')
    )
    WITH CHECK (
        current_setting('app.current_role', true) IN ('admin', 'editor')
    );

-- Units:
-- admin/editor: all units, including unassigned units
-- tenant: units belonging to buildings assigned to current tenant
-- public: no direct unit access

CREATE POLICY units_select_policy ON units
    FOR SELECT
    USING (
        current_setting('app.current_role', true) IN ('admin', 'editor')
        OR (
            current_setting('app.current_role', true) = 'tenant'
            AND EXISTS (
                SELECT 1
                FROM buildings b
                WHERE b.id = units.building_id
                  AND b.tenant_id = NULLIF(
                      current_setting('app.current_tenant_id', true), ''
                  )::UUID
                  AND b.deleted_at IS NULL
            )
        )
    );

CREATE POLICY units_modify_policy ON units
    FOR ALL
    USING (
        current_setting('app.current_role', true) IN ('admin', 'editor')
    )
    WITH CHECK (
        current_setting('app.current_role', true) IN ('admin', 'editor')
    );

-- =========================================================
-- 5. Amenities:
-- admin/editor: all amenities
-- tenant: global amenities and own tenant's amenities
-- public: no access
-- =========================================================

ALTER TABLE amenities ENABLE ROW LEVEL SECURITY;

DROP POLICY IF EXISTS amenities_access_policy ON amenities;
DROP POLICY IF EXISTS amenities_select_policy ON amenities;
DROP POLICY IF EXISTS amenities_modify_policy ON amenities;

CREATE POLICY amenities_select_policy ON amenities
    FOR SELECT
    USING (
        current_setting('app.current_role', true) IN ('admin', 'editor')
        OR (
            current_setting('app.current_role', true) = 'tenant'
            AND (
                is_global = TRUE
                OR tenant_id = NULLIF(
                    current_setting('app.current_tenant_id', true), ''
                )::UUID
            )
        )
    );

CREATE POLICY amenities_modify_policy ON amenities
    FOR ALL
    USING (
        current_setting('app.current_role', true) IN ('admin', 'editor')
    )
    WITH CHECK (
        current_setting('app.current_role', true) IN ('admin', 'editor')
    );

-- =========================================================
-- 6. Media:
-- admin/editor: all media
-- tenant: media belonging to own tenant
-- public: no direct access
-- =========================================================

ALTER TABLE media ENABLE ROW LEVEL SECURITY;

DROP POLICY IF EXISTS media_select_policy ON media;
DROP POLICY IF EXISTS media_modify_policy ON media;

CREATE POLICY media_select_policy ON media
    FOR SELECT
    USING (
        current_setting('app.current_role', true) IN ('admin', 'editor')
        OR (
            current_setting('app.current_role', true) = 'tenant'
            AND tenant_id = NULLIF(
                current_setting('app.current_tenant_id', true), ''
            )::UUID
        )
    );

CREATE POLICY media_modify_policy ON media
    FOR ALL
    USING (
        current_setting('app.current_role', true) IN ('admin', 'editor')
    )
    WITH CHECK (
        current_setting('app.current_role', true) IN ('admin', 'editor')
    );

COMMIT;
