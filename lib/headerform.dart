import 'dart:html';

import 'package:mango_ui/bodies/header.dart';
import 'package:mango_ui/formstate.dart';

import 'package:mango_ui/trustvalidator.dart';
import 'models/headeritem.dart';

class HeaderForm extends FormState {
  List<Header> _items;

  HeaderForm(String idElem, String btnSubmit, String btnAdd)
      : super(idElem, btnSubmit) {
    _items = findHeaders();

    querySelector(btnAdd).onClick.listen(onAddClick);
  }

  void onAddClick(MouseEvent e) {
    addHeader();
  }

  List<Header> get items {
    return _items;
  }

  void addHeader() {
    //add elements
    var indx = items.length;

    var schema = buildElement(indx);
    form.children.add(schema);

    var item = new HeaderItem("#uplHeaderImg${indx}",
        "#txtHeaderHeading${indx}", "#txtHeaderText${indx}");

    _items.add(item.toDTO());
  }

  List<Header> findHeaders() {
    var hasHeaders = false;
    var result = new List<Header>();
    var indx = 0;

    do {
      var headr = new HeaderItem("#uplHeaderImg${indx}",
          "#txtHeaderHeading${indx}", "#txtHeaderText${indx}");

      hasHeaders = headr.loaded();

      if (hasHeaders) {
        result.add(headr.toDTO());
      }

      indx++;
    } while (hasHeaders);

    return result;
  }

  //returns HTML for a Header Item
  Element buildElement(int index) {
    var schema = '''
    <div class="card">
            <header class="card-header">
                <p class="card-header-title">
                    <div class="control">
                        <input type="text" class="input" id="txtHeaderHeading${index}" required placeholder="Heading"
                            value="New Heading" />
                        <p class="help is-danger"></p>
                    </div>
                </p>
                <a href="#" data-liID="liHeader${index}" class="card-header-icon" aria-label="more options">
                    <span class="icon">
                        <i class="fa fa-close" aria-hidden="true"></i>
                    </span>
                </a>
            </header>
            <div class="card-image">
                <figure class="image">
                    <img id="uplHeaderView${index}" class="is-hidden" src="" alt="unknown image">
                </figure>
            </div>
            <div class="card-content">
                <div class="content">
                    <div class="field">
                        <div class="control">
                            <label class="label" for="uplHeaderImg${index}">Banner</label>
                            <div class="file">
                                <label class="file-label">
                                    <input class="file-input" type="file" multiple="false" data-for="banner"
                                        data-name="Header" data-id="0`0" accept=".jpg, .jpeg, .png"
                                        id="uplHeaderImg${index}" placeholder="Header Banner" required />
                                    <p class="help is-danger"></p>
                                    <span class="file-cta">
                                        <span class="file-icon">
                                            <i class="fas fa-upload"></i>
                                        </span>
                                        <span class="file-label">
                                            Choose a file…
                                        </span>
                                    </span>
                                </label>
                            </div>
                        </div>
                    </div>
                    <div class="field">
                        <div class="control">
                            <label class="label" for="txtHeaderText${index}">Text</label>
                            <textarea id="txtHeaderText${index}"
                            class="textarea"  placeholder="Text" required cols="40" rows="5"></textarea>
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
