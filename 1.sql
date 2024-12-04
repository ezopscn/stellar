SELECT
    CASE 
        WHEN vt2.b_pre_bw BETWEEN 0 AND 20 THEN '0-20'
        WHEN vt2.b_pre_bw BETWEEN 20 AND 30 THEN '20-30'
        WHEN vt2.b_pre_bw BETWEEN 30 AND 50 THEN '30-50'
        WHEN vt2.b_pre_bw BETWEEN 50 AND 100 THEN '50-100'
        ELSE '100+'
    END AS bandwidth,
    count(vt1.a_sn),
    sum(vt1.a_avg_bw),
    sum(vt2.b_pre_bw),
    sum(vt2.b_eve_peak95_bw)
FROM
    (
        SELECT
            DISTINCT a.sn as a_sn,
            a.avg_bw as a_avg_bw,
            a.device_type2 as a_device_type2
        FROM
            iaas_app.app_d1_iaas_docker_sn_kc_info_v1 a
        WHERE
            a.stime >= yesterday()
            AND a.device_type2 = 'oedgestation'
    ) as vt1
    LEFT JOIN (
        SELECT
            DISTINCT b.sn as b_sn,
            b.pre_bw * 1000 as b_pre_bw,
            b.eve_peak95_bw * 1000 as b_eve_peak95_bw
        FROM
            iaas_app.app_d1_iaas_biz_sn_info b
        where
            b.stime = yesterday()
) as vt2 ON vt1.a_sn = vt2.b_sn;