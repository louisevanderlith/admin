import 'dart:html';

import '../formstate.dart';
import '../trustvalidator.dart';
import 'socialmediaitem.dart';

class SocialmediaForm extends FormState {
  List<SocialmediaItem> _items;

  SocialmediaForm(String idElem, String btnSubmit, String btnAdd)
      : super(idElem, btnSubmit) {
    _items = findSocials();

    querySelector(btnAdd).onClick.listen(onAddClick);
  }

  void onAddClick(MouseEvent e) {
    addSocial();
  }

  List<SocialmediaItem> get items {
    return _items;
  }

  List<SocialmediaItem> findSocials() {
    var hasSocial = false;
    var result = new List<SocialmediaItem>();
    var indx = 0;

    do {
      var social =
          new SocialmediaItem('#txtSocialIcon${indx}', '#txtSocialURL${indx}');
      hasSocial = social.loaded();

      if (hasSocial) {
        result.add(social);
      }

      indx++;
    } while (hasSocial);

    return result;
  }

  void addSocial() {
    //add elements
    var indx = items.length;

    var schema = buildElement(indx);
    form.children.add(schema);
  
    var item =
        new SocialmediaItem('#txtSocialIcon${indx}', '#txtSocialURL${indx}');

    _items.add(item);
  }

  //returns HTML for a SocialMedia Item
  Element buildElement(int index) {
    var schema = '''
    <div class="card">
                <header class="card-header">
                    <a href="#" data-liID="liSocial${index}" class="card-header-icon" aria-label="more options">
                        <span class="icon">
                            <i class="fas fa-close" aria-hidden="true"></i>
                        </span>
                    </a>
                </header>
                <div class="card-content">
                    <div class="content">
                        <div class="field">
                            <div class="control">
                                <label class="label" for="txtSocialURL${index}">URL
                                    <input class="input" type="text" min-length="3" id="txtSocialURL${index}" required
                                        value="" />
                                    <p class="help is-danger"></p>
                            </div>
                        </div>
                        <div class="field">
                            <div class="control">
                                <label class="label" for="txtSocialIcon${index}">Icon
                                    <input class="input" type="text" id="txtSocialIcon${index}" required
                                        value="" />
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
