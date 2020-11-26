import 'dart:html';

import 'package:Admin.APP/bodies/categorystock.dart';
import 'package:mango_stock/bodies/stockitem.dart';
import 'package:mango_ui/keys.dart';
import 'package:mango_ui/trustvalidator.dart';

class CategoryStockForm {
  DivElement form;
  final Key categoryKey;

  CategoryStockForm(this.categoryKey) {
    form = querySelector("#dvStock");
    registerEvents();
  }

  void registerEvents() {
    querySelector("#btnAddItem").onClick.listen(onAddClick);

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
      var item = new CategoryStock(
        "#cboItems${indx}",
        "#txtShortName${indx}",
        "#uplThumbImg${indx}",
        "#hdnOwnerKey${indx}",
        "#txtExpires${indx}",
        "#txtCurrency${indx}",
        "#numPrice${indx}",
        "#numEstimate${indx}",
        "#lstTags${indx}",
        "#txtLocation${indx}",
        "#numViews${indx}",
        "#lstHistory${indx}",
      );

      isLoaded = item.loaded;

      if (isLoaded) {
        result.add(item.toDTO());
      }

      indx++;
    } while (isLoaded);

    return result;
  }

  void addItem() {
    var schema = buildElement(items.length);
    form.children.add(schema);
    registerEvents();
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
                            <input class="input" id="txtExpires${index}" type="datetime-local" required
                                   value=""/>
                            <p class="help is-danger"></p>
                        </div>
                    </div>
                    <div class="field">
                        <label for="txtCurrency${index}" class="label">Expiry Date:</label>
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
