{{ config(materialized='external') }}

with

source as (
    select * from {{ source('conf', 'raw_attendees') }}
),

renamed as (

    select
        attendee_id as participant_id,
        email,
        country,
        ticket_type,
        first_name || ' ' || last_name AS full_name
    from source

)

select * from renamed