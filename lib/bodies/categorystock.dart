import 'dart:html';

import 'package:mango_stock/bodies/stockitem.dart';
import 'package:mango_ui/keys.dart';

class CategoryStock {
  HiddenInputElement hdnItemKey;
  TextInputElement txtShortName;
  FileUploadInputElement uplImage;
  HiddenInputElement hdnOwnerKey;
  DateInputElement txtExpires;
  TextInputElement txtCurrency;
  NumberInputElement numPrice;
  NumberInputElement numEstimate;
  UListElement lstTags;
  TextInputElement txtLocation;
  UListElement lstHistory;
  NumberInputElement numViews;

  bool _loaded;

  CategoryStock(
      String itemKeyId,
      String shortnameId,
      String imageId,
      String ownerId,
      String expiresId,
      String currencyId,
      String priceId,
      String estId,
      String tagsId,
      String locationId,
      String viewsId,
      String historyId) {
    hdnItemKey = querySelector(itemKeyId);
    txtShortName = querySelector(shortnameId);
    uplImage = querySelector(imageId);
    hdnOwnerKey = querySelector(ownerId);
    txtExpires = querySelector(expiresId);
    txtCurrency = querySelector(currencyId);
    numPrice = querySelector(priceId);
    numEstimate = querySelector(estId);
    lstTags = querySelector(tagsId);
    txtLocation = querySelector(locationId);
    lstHistory = querySelector(historyId);
    numViews = querySelector(viewsId);

    _loaded = hdnItemKey != null &&
        uplImage != null &&
        hdnOwnerKey != null &&
        txtExpires != null &&
        txtCurrency != null &&
        numPrice != null &&
        lstTags != null &&
        txtLocation != null &&
        lstHistory != null &&
        numViews != null;
  }

  bool get loaded {
    return _loaded;
  }

  Key get itemKey {
    return new Key(hdnItemKey.value);
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
    return lstTags.children.map((e) => e.text);
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

  StockItem toDTO() {
    return new StockItem(itemKey, shortName, imageKey, ownerKey, expires,
        currency, price, estimate, tags, location, views, history);
  }
}
