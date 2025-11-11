{{ config(materialized='external') }}

WITH
  attendees as (
    select * from {{ ref('stg_attendees') }}
)

select
    country,
    count(*) as citizens
from attendees
group by country
order by citizens desc