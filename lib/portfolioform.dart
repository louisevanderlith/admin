import 'dart:html';

import 'package:mango_ui/bodies/portfolio.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/trustvalidator.dart';

import 'models/portfolioitem.dart';

class PortfolioForm extends FormState {
  PortfolioForm(String idElem, String btnSubmit, String btnAdd)
      : super(idElem, btnSubmit) {
        findPortfolios();
    querySelector(btnAdd).onClick.listen(onAddClick);
  }

  void onAddClick(MouseEvent e) {
    addPortfolio();
  }

  List<Portfolio> get items {
    return findPortfolios();
  }
  
  List<Portfolio> findPortfolios() {
    var hasPortf = false;
    var result = new List<Portfolio>();
    var indx = 0;

    do {
      var portf = new PortfolioItem('#uplPortfolioImg${indx}',
          '#txtPortfolioName${indx}', '#txtPortfolioURL${indx}');

      hasPortf = portf.loaded();

      if (hasPortf) {
        result.add(portf.toDTO());
      }

      indx++;
    } while (hasPortf);

    return result;
  }

  void addPortfolio() {
    //add elements
    var indx = items.length;

    var schema = buildElement(indx);
    form.children.add(schema);

    findPortfolios();
  }

  //returns HTML for a Portfolio Item
  Element buildElement(int index) {
    var schema = '''
      <div class="card">
            <header class="card-header">
                <a href="#" data-liID="liPortfolio${index}" class="card-header-icon" aria-label="more options">
                    <span class="icon">
                        <i class="fa fa-close" aria-hidden="true"></i>
                    </span>
                </a>
            </header>
            <div class="card-image">
                <figure class="image">
                    <img id="uplPortfolioView${index}" src="" is-hidden alt="">
                </figure>
            </div>
            <div class="card-content">
                <div class="content">
                    <div class="field">
                        <div class="control">
                            <label class="label" for="uplPortfolioImg${index}">Portfolio image</label>
                            <div class="file">
                                <label class="file-label">
                                    <input class="file-input" type="file" multiple="false" data-for="thumb"
                                        data-name="Portfolio" data-id="" accept=".jpg, .jpeg, .png"
                                        id="uplPortfolioImg${index}" placeholder="Portfolio image"
                                        required />
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
                            <label class="label" for="txtPortfolioName${index}">Name</label>
                            <input type="text" class="input" id="txtPortfolioName${index}" required placeholder="Name"
                                value="" />
                            <p class="help is-danger"></p>
                        </div>
                    </div>
                    <div class="field">
                        <div class="control">
                            <label class="label" for="txtPortfolioURL${index}">URL</label>
                            <input class="input" type="url" id="txtPortfolioURL${index}" required placeholder="URL"
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
