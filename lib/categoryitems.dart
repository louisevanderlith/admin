import 'dart:html';

import 'package:Admin.APP/components/categorystockitem.dart';
import 'package:mango_stock/bodies/stockitem.dart';
import 'package:mango_ui/keys.dart';
import 'package:mango_ui/trustvalidator.dart';

import 'components/categorystockitem.dart';

class CategoryItemsForm {
  DivElement form;
  final Key categoryKey;
  List<CategoryStockItem> stock;

  CategoryItemsForm(this.categoryKey) {
    form = querySelector("#dvStock");
    querySelector("#btnAddItem").onClick.listen(onAddClick);

    stock = findItems();
  }

  void onAddClick(MouseEvent e) {
    addItem();
  }

  List<StockItem> get items {
    return stock.map((e) => e.toDTO()).toList(growable: false);
  }

  List<StockItem> get simpleitems {
    final result = new List<StockItem>();

    for (var item in stock) {
      result.add(item.toDTO());
    }

    return result;
  }

  List<CategoryStockItem> findItems() {
    var isLoaded = false;
    var result = new List<CategoryStockItem>();
    var indx = 0;

    do {
      try {
        var item = new CategoryStockItem(
          "#cboItems${indx}",
          "#txtShortName${indx}",
          "#uplThumbImg${indx}",
          "#hdnOwnerKey${indx}",
          "#hdnExpires${indx}",
          "#txtExpires${indx}",
          "#txtCurrency${indx}",
          "#numPrice${indx}",
          "#numEstimate${indx}",
          "#lstTags${indx}",
          "#txtLocation${indx}",
          "#numViews${indx}",
          "#lstHistory${indx}",
          "#numQuantity${indx}",
        );

        isLoaded = item.loaded;

        if (isLoaded) {
          result.add(item);
        }
      } catch (exc) {
        print(exc);
      }

      indx++;
    } while (isLoaded);

    return result;
  }

  void addItem() {
    final index = items.length;
    var schema = buildElement(index);
    form.children.add(schema);



    stock = findItems();
  }

  void populateItemList(num index, String category) {
      //Add items to list
  }

  //returns HTML for this Item
  Element buildElement(int index) {
    var schema = '''
    <div class="card">
            <header class="card-header">
                <p class="card-header-title">
                    New Item
                </p>
                <a class="card-header-icon" aria-label="more options">
                <span class="icon">
                    <i class="fas fa-angle-down" aria-hidden="true"></i>
                </span>
                </a>
            </header>
            <div class="card-content" hidden>
                <div class="content">
                    <div class="field">
                        <label class="label">Item</label>
                        <div class="control">
                            <select class="select">
                            </select>
                        </div>
                    </div>
                    <input type="hidden" id="hdnOwnerKey${index}">
                    <div class="field">
                        <label class="label">Short Name:</label>
                        <div class="control">
                            <input class="input" id="txtShortName${index}" placeholder="Name" type="text"
                                   min-length="3"
                                   required/>
                            <p class="help is-danger"></p>
                        </div>
                    </div>
                    <div class="field">
                        <label class="label" for="uplThumbImg">Thumbnail</label>
                        <div class="control">
                            <div class="file">
                                <label class="file-label">
                                    <input class="file-input" type="file" multiple="false" data-for="thumb"
                                           data-name="Banner"
                                           data-id="" accept=".jpg, .jpeg, .png"
                                           id="uplThumbImg"
                                           placeholder="Stock Thumbnail" require />
                                    <p class="help is-danger"></p>
                                    <span class="file-cta">
                                        <span class="file-icon">
                                            <i class="fas fa-upload"></i>
                                        </span>
                                        <span class="file-label">
                                            Choose an image.
                                        </span>
                                    </span>
                                </label>
                            </div>
                            <input type="image" id="uplThumbView"
                                   class='image is-hidden'
                                   src=""
                                   alt="profile image"/>
                            <p class="help is-danger"></p>
                        </div>
                    </div>
                    <div class="field">
                        <label class="label">Price:</label>
                        <div class="control">
                            <input class="input" id="txtItemPrice${index}" type="number" min="1" required
                                   value=""/>
                            <p class="help is-danger"></p>
                        </div>
                    </div>
                    <div class="field">
                        <label for="txtExpires${index}" class="label">Expiry Date:</label>
                        <div class="control">
                            <input class="input" id="txtExpires${index}" type="date" required
                                   value=""/>
                            <p class="help is-danger"></p>
                        </div>
                    </div>
                    <div class="field">
                        <label for="txtCurrency${index}" class="label">Currency:</label>
                        <div class="control">
                            <input class="input" id="txtCurrency${index}" type="text" required
                                   value=""/>
                            <p class="help is-danger"></p>
                        </div>
                    </div>
                    <div class="field">
                        <label for="numPrice${index}" class="label">Price:</label>
                        <div class="control">
                            <input class="input" id="numPrice${index}" type="number" required
                                   value=""/>
                            <p class="help is-danger"></p>
                        </div>
                    </div>
                    <div class="field">
                        <label for="numEstimate${index}" class="label">Estimate Value:</label>
                        <div class="control">
                            <input class="input" id="numEstimate${index}" type="number" required
                                   value=""/>
                            <p class="help is-danger"></p>
                        </div>
                    </div>
                    <div class="field">
                        <label for="lstTags${index}" class="label">Tags:</label>
                        <div class="control">
                            <ul id="lstTags${index}">
                            </ul>
                            <p class="help is-danger"></p>
                        </div>
                    </div>
                    <div class="field">
                        <label for="txtLocation${index}" class="label">Location:</label>
                        <div class="control">
                            <input class="input" id="txtLocation${index}" type="text" required
                                   value=""/>
                            <p class="help is-danger"></p>
                        </div>
                    </div>
                    <div class="field">
                        <label for="numViews${index}" class="label">Views:</label>
                        <div class="control">
                            <input class="input" id="numViews${index}" disabled type="number" required
                                   value=""/>
                            <p class="help is-danger"></p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        ''';

    return new Element.html(schema, validator: new TrustedNodeValidator());
  }
}
