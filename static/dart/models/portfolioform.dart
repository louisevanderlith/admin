import '../formstate.dart';
import 'portfolioitem.dart';

class PortfolioForm extends FormState {
  List<PortfolioItem> _items;

  PortfolioForm(String idElem, btnSubmit) : super(idElem, btnSubmit) {
    _items = findPortfolios(btnSubmit);
  }

  List<PortfolioItem> get items {
    return _items;
  }

  List<PortfolioItem> findPortfolios(String btnSubmit) {
    var hasPortf = false;
    var result = new List<PortfolioItem>();
    var indx = 0;

    do {
      var portf = new PortfolioItem('#uplPortfolioImg${indx}',
          '#txtPortfolioName${indx}', '#txtPortfolioURL${indx}');

      hasPortf = portf.loaded();

      if (hasPortf) {
        result.add(portf);
      }

      indx++;
    } while (hasPortf);

    return result;
  }
}
