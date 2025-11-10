{{ config(materialized='external') }}

with

source as (
    select * from {{ source('conf', 'raw_sessions') }}
),

renamed as (

    select
        Id as session_id,
        Title,
        Track,
        Start Time,
        Description,
        Speakers
    from source

)

select * from renamed