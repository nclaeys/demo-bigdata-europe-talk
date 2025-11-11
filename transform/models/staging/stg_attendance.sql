{{ config(materialized='external') }}

with

source as (
    select * from {{ source('conf', 'attendance') }}
),

renamed as (
    select
        attendee_id as participant_id,
        session_id,
        start_time,
    from source

)

select * from renamed