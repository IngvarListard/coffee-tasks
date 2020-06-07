SELECT id,
       name
  FROM
      (SELECT id,
              name,

              (SELECT COUNT(*)
                 FROM tags_goods
                WHERE goods_id=id ) AS t
         FROM goods
        WHERE t=
              (SELECT COUNT(*)
                 FROM tags))
