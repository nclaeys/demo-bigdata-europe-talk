{{ config(materialized='external') }}

with

source as (
    select * from {{ source('conf', 'raw_feedback') }}
),

renamed as (

    select
        feedback_id,
        session_id,
        attendee_id,
        rating,
        comments,
    from source

)

select * from renamed