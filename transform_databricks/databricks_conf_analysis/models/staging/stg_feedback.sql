with

source as (
    select * from {{ source('conf', 'feedback') }}
),

renamed as (

    select
        feedback_id,
        session_id,
        attendee_id,
        rating,
        comments
    from source

)

select * from renamed