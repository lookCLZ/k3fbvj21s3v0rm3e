SELECT m_prod_catalogues.cata_name,merchants.merchant_name
,m_products.* FROM `m_products` LEFT JOIN m_prod_catalogues
 ON m_products.cata_id = m_prod_catalogues.id LEFT JOIN
  merchants ON m_products.merchant_id =
   merchants.id WHERE `m_products`.`deleted_at`
	IS NULL ORDER BY m_products.created_at desc
	 LIMIT 20 OFFSET 0 





///////////////

SELECT m_prod_catalogues.cata_name,merchants.merchant_name
,m_products.* FROM `m_products` LEFT JOIN m_prod_catalogues
 ON m_products.cata_id = m_prod_catalogues.id LEFT JOIN 
 merchants ON m_products.merchant_id = merchants.id LEFT
  JOIN m_good_item_goods ON m_products.id =
   m_good_item_goods.prod_id WHERE `m_products`
   .`deleted_at` IS NULL ORDER BY m_products.created_at
    desc LIMIT 20 OFFSET 0  