WITH
  sessions as (
    select * from {{ ref('stg_sessions') }}
),
    feedback as (
        select * from {{ ref('stg_feedback') }}
    )

SELECT
    AVG(f.rating) as avg_rating,
    MAX(f.rating) as max_rating,
    MIN(f.rating) as min_rating,
    track
from sessions s
         join feedback f on f.session_id = s.session_id
where s.track is not null
group by s.track
