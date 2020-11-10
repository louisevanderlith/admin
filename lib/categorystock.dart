import 'package:Admin.APP/bodies/categorystock.dart';
import 'package:mango_stock/bodies/stockitem.dart';

class CategoryStockForm {
  DivElement form;

  ContactsForm() {
    form = querySelector("#dvItems");
    querySelector("#btnAddStock").onClick.listen(onAddClick);
  }

  void onAddClick(MouseEvent e) {
    addItem();
  }

  List<StockItem> get items {
    return findItems();
  }

  List<StockItem> findItems() {
    var isLoaded = false;
    var result = new List<StockItem>();
    var indx = 0;

    do {
      var item = new CategoryStock("#hdnItemKey${indx}",
          "#txtShortName${indx}",
          "#upl"
          "#txtContactValue${indx}");

      isLoaded = item.loaded;

      if (isLoaded) {
        result.add(item.toDTO());
      }

      indx++;
    } while (isLoaded);

    return result;
  }
}
