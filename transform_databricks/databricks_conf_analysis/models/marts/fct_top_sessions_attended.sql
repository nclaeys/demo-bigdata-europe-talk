WITH
  sessions as (
    select * from {{ ref('stg_sessions') }}
  ),
  attendance as (
    select * from {{ ref('stg_attendance') }}
  )

SELECT
    s.title,
    count(*) as attendance_count
from sessions s
         join attendance a on a.session_id = s.session_id
where s.track is not null
group by s.title
order by attendance_count desc
limit 10
