with

source as (
    select * from {{ source('conf', 'sessions') }}
),

renamed as (

    select
        Id as session_id,
        Title as title,
        Track as track,
        "Start Time" as start_time,
        Description as description,
        Speakers as speakers
    from source

)

select * from renamed