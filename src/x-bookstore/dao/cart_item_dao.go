package dao

import (
	"x-bookstore/model"
	"x-bookstore/utils"
)

//AddCartItem 向购物项表中插入购物项
func AddCartItem(cartItem *model.CartItem) error {
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"

	_, err := utils.DB.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

//GetCartItemByBookIDAndCartID 根据图书的id和购物车的id获取对应的购物项
func GetCartItemByBookIDAndCartID(bookID string, cartID string) (*model.CartItem, error) {
	//sql
	sqlStr := "select id,count,amount,cart_id,ctime,mtime from cart_items where book_id = ? and cart_id = ?"
	row := utils.DB.QueryRow(sqlStr, bookID, cartID)

	var cartItem = &model.CartItem{}
	err := row.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &cartItem.CartID, &cartItem.Ctime,&cartItem.Mtime)
	if err != nil {
		return nil, err
	}

	book, _ := GetBookByID(bookID)
	//根据图书的id查询图书信息
	//将book设置到购物项
	cartItem.Book = book
	return cartItem, nil
}

//GetCartItemsByCartID 根据购物车的id获取购物车中所有的购物项
func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {
	//写sql语句
	sqlStr := "select id,count,amount,book_id,cart_id from cart_items where cart_id = ?"
	//执行
	rows, err := utils.DB.Query(sqlStr, cartID)
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		//设置一个变量接收bookId
		var bookID string
		//创建cartItem
		cartItem := &model.CartItem{}
		err2 := rows.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &bookID, &cartItem.CartID)
		if err2 != nil {
			return nil, err2
		}
		//根据bookID获取图书信息
		book, _ := GetBookByID(bookID)
		//将book设置到购物项中
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

//UpdateBookCount 根据购物项中的相关信息更新购物项中图书的数量和金额小计
func UpdateBookCount(cartItem *model.CartItem) error {
	//写sql语句
	sql := "update cart_items set count = ? , amount = ? where book_id = ? and cart_id = ?"
	//执行
	_, err := utils.DB.Exec(sql, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteCartItemsByCartID 根据购物车的id删除所有的购物项
func DeleteCartItemsByCartID(cartID string) error {
	//写sql语句
	sql := "delete from cart_items where cart_id = ?"
	_, err := utils.DB.Exec(sql, cartID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteCartItemByID 根据购物项的id删除购物项
func DeleteCartItemByID(cartItemID string) error {
	//写sql语句
	sql := "delete from cart_items where id = ?"
	//执行
	_, err := utils.DB.Exec(sql, cartItemID)
	if err != nil {
		return err
	}
	return nil
}