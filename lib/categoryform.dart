import 'dart:html';

import 'package:dart_toast/dart_toast.dart';
import 'package:mango_stock/bodies/category.dart';
import 'package:mango_stock/stockapi.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/keys.dart';

import 'categoryinfo.dart';
import 'categorystock.dart';

class CategoryForm extends FormState {
  Key objKey;

  CategoryInfoForm info;
  CategoryStockForm stock;

  CategoryForm(Key k) : super("#frmCategory", "#btnSubmit") {
    objKey = k;

    info = new CategoryInfoForm();
    stock = new CategoryStockForm();

    querySelector("#btnSubmit").onClick.listen(onSubmitClick);
  }

  void onSubmitClick(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);

      final obj = new Category(
          info.client, info.text, info.description, info.image, stock.items);

      HttpRequest req;
      if (objKey.toJson() != "0`0") {
        req = await updateCategory(objKey, obj);
        if (req.status == 200) {
          Toast.success(
              title: "Success!",
              message: req.response,
              position: ToastPos.bottomLeft);
        } else {
          Toast.error(
              title: "Failed!",
              message: req.response,
              position: ToastPos.bottomLeft);
        }
      } else {
        req = await createCategory(obj);

        if (req.status == 200) {
          final key = req.response;
          objKey = new Key(key);

          Toast.success(
              title: "Success!",
              message: "Content Saved",
              position: ToastPos.bottomLeft);
        } else {
          Toast.error(
              title: "Failed!",
              message: req.response,
              position: ToastPos.bottomLeft);
        }
      }
    }
  }
}
