import 'dart:html';

import 'package:mango_artifact/uploadapi.dart';
import 'package:mango_stock/bodies/stockitem.dart';
import 'package:mango_ui/keys.dart';

class CategoryStockItem {
  SelectElement cboItems;
  TextInputElement txtShortName;
  FileUploadInputElement uplImage;
  HiddenInputElement hdnOwnerKey;
  HiddenInputElement hdnExpires;
  DateInputElement txtExpires;
  TextInputElement txtCurrency;
  NumberInputElement numPrice;
  NumberInputElement numEstimate;
  UListElement lstTags;
  TextInputElement txtLocation;
  UListElement lstHistory;
  NumberInputElement numViews;
  NumberInputElement numQuantity;

  bool _loaded;

  CategoryStockItem(
      String itemsId,
      String shortnameId,
      String imageId,
      String ownerId,
      String expiresHdnId,
      String expiresId,
      String currencyId,
      String priceId,
      String estId,
      String tagsId,
      String locationId,
      String viewsId,
      String historyId,
      String quantityId) {
    cboItems = querySelector(itemsId);
    txtShortName = querySelector(shortnameId);
    uplImage = querySelector(imageId);
    hdnOwnerKey = querySelector(ownerId);
    hdnExpires = querySelector(expiresHdnId);
    txtExpires = querySelector(expiresId);
    txtCurrency = querySelector(currencyId);
    numPrice = querySelector(priceId);
    numEstimate = querySelector(estId);
    lstTags = querySelector(tagsId);
    txtLocation = querySelector(locationId);
    lstHistory = querySelector(historyId);
    numViews = querySelector(viewsId);
    numQuantity = querySelector(quantityId);

    if (uplImage != null) {
      uplImage.onChange.listen(uploadFile);
    }

    _loaded = cboItems != null &&
        uplImage != null &&
        hdnOwnerKey != null &&
        hdnExpires != null &&
        txtExpires != null &&
        txtCurrency != null &&
        numPrice != null &&
        lstTags != null &&
        txtLocation != null &&
        lstHistory != null &&
        numViews != null;

    txtExpires.value = hdnExpires.value.replaceAll(" 00:00:00 +0000 UTC", "");
  }

  bool get loaded {
    return _loaded;
  }

  Key get itemKey {
    return new Key(cboItems.value);
  }

  String get shortName {
    return txtShortName.value;
  }

  Key get imageKey {
    return new Key(uplImage.dataset["id"]);
  }

  Key get ownerKey {
    return new Key(hdnOwnerKey.value);
  }

  DateTime get expires {
    return txtExpires.valueAsDate;
  }

  String get currency {
    return txtCurrency.value;
  }

  double get price {
    return numPrice.valueAsNumber;
  }

  double get estimate {
    return numEstimate.valueAsNumber;
  }

  List<String> get tags {
    return lstTags.children.map((e) => e.innerText).toList();
  }

  String get location {
    return txtLocation.value;
  }

  num get views {
    return numViews.valueAsNumber;
  }

  Map<DateTime, Key> get history {
    return new Map<DateTime, Key>();
  }

  num get quantity {
    return numQuantity.valueAsNumber;
  }

  StockItem toDTO() {
    return new StockItem(itemKey, shortName, imageKey, ownerKey, expires,
        currency, price, estimate, tags, location, views, history, quantity);
  }
}
