from importlib import reload
from typing import TYPE_CHECKING

try:
    import pandas as pd
    from pandas.api.types import is_numeric_dtype
except ImportError:
    pass

from hstest.testing.plotting.drawing.drawing import Drawing
from hstest.testing.plotting.drawing.drawing_builder import DrawingBuilder
from hstest.testing.plotting.drawing.drawing_library import DrawingLibrary
from hstest.testing.plotting.drawing.drawing_type import DrawingType
from hstest.testing.plotting.matplotlib_handler import MatplotlibHandler

if TYPE_CHECKING:
    from hstest.testing.runner.plot_testing_runner import DrawingsStorage


class SeabornHandler:
    _saved = False
    _replaced = False

    _displot = None
    _histplot = None
    _lineplot = None
    _lmplot = None
    _scatterplot = None
    _catplot = None
    _barplot = None
    _violinplot = None
    _heatmap = None
    _boxplot = None

    @staticmethod
    def replace_plots(drawings: 'DrawingsStorage'):
        try:
            import seaborn as sns
            import numpy as np
        except ModuleNotFoundError:
            return

        def displot(data=None, **kwargs):
            x = None if 'x' not in kwargs else kwargs['x']
            y = None if 'y' not in kwargs else kwargs['y']

            if data is None:
                curr_data = {
                    'x': np.array(x),
                    'y': np.array(y)
                }

                drawing = Drawing(
                    DrawingLibrary.seaborn,
                    DrawingType.dis,
                    None,
                    kwargs,
                )
                drawings.append(drawing)
                return

            x_arr = []
            y_arr = []

            if not x and not y:
                for column in data.columns:
                    if not is_numeric_dtype(data[column]):
                        continue

                    curr_data = {
                        'x': np.array([column]),
                        'y': data[column].to_numpy()
                    }

                    drawing = Drawing(
                        DrawingLibrary.seaborn,
                        DrawingType.dis,
                        None,
                        kwargs,
                    )
                    drawings.append(drawing)
                return

            if x:
                x_arr = data[x].to_numpy()
            if y:
                y_arr = data[y].to_numpy()

            curr_data =  {
                'x': np.array(x_arr),
                'y': np.array(y_arr)
            }

            drawing = Drawing(
                DrawingLibrary.seaborn,
                DrawingType.dis,
                None,
                kwargs,
            )
            drawings.append(drawing)

        def histplot(data=None, **kwargs):
            if data is not None:
                if 'y' in kwargs and kwargs['y'] is not None:
                    drawings.append(
                        DrawingBuilder.get_hist_drawing(
                            data[kwargs['y']],
                            DrawingLibrary.seaborn,
                            kwargs,
                        )
                    )
                elif 'x' in kwargs and kwargs['x'] is not None:
                    drawings.append(
                        DrawingBuilder.get_hist_drawing(
                            data[kwargs['x']],
                            DrawingLibrary.seaborn,
                            kwargs,
                        )
                    )
                else:
                    if type(data) == pd.DataFrame:
                        for col in data.columns:
                            drawings.append(
                                DrawingBuilder.get_hist_drawing(
                                    data[col],
                                    DrawingLibrary.seaborn,
                                    kwargs,
                                )
                            )
                    else:
                        drawings.append(
                            DrawingBuilder.get_hist_drawing(
                                data,
                                DrawingLibrary.seaborn,
                                kwargs,
                            )
                        )

        def lineplot(*, data=None, x=None, y=None, **kwargs):
            if x is not None:
                x_array = data[x].to_numpy()
            else:
                x_array = data.index.to_numpy()

            if y is not None:
                y_array = data[y].to_numpy()

                drawings.append(
                    DrawingBuilder.get_line_drawing(
                        x_array,
                        y_array,
                        DrawingLibrary.seaborn,
                        kwargs,
                    )
                )
                return drawings

            for column in data.columns:
                if not is_numeric_dtype(data[column]):
                    continue

                drawings.append(
                    DrawingBuilder.get_line_drawing(
                        x_array,
                        data[column],
                        DrawingLibrary.seaborn,
                        kwargs,
                    )
                )

        def lmplot(x=None, y=None, data=None, **kwargs):
            curr_data = {
                'data': data,
                'x': x,
                'y': y,
                'kwargs': kwargs
            }

            drawing = Drawing(
                DrawingLibrary.seaborn,
                DrawingType.lm,
                None,
                kwargs,
            )
            drawings.append(drawing)

        def scatterplot(x=None, y=None, data=None, **kwargs):
            if x is not None and y is not None:
                drawings.append(
                    DrawingBuilder.get_scatter_drawing(
                        data[x], data[y],
                        DrawingLibrary.seaborn,
                        kwargs,
                    )
                )
                return

            if x is None and y is None and data is not None:
                for column in data.columns:
                    if not is_numeric_dtype(data[column]):
                        continue

                    x = data.index
                    drawings.append(
                        DrawingBuilder.get_scatter_drawing(
                            x, data[column],
                            DrawingLibrary.seaborn,
                            kwargs,
                        )
                    )

        def catplot(x=None, y=None, data=None, **kwargs):
            curr_data = {
                'data': data,
                'x': x,
                'y': y,
                'kwargs': kwargs
            }

            drawing = Drawing(
                DrawingLibrary.seaborn,
                DrawingType.cat,
                None,
                kwargs,
            )
            drawings.append(drawing)

        def barplot(x=None, y=None, data=None, **kwargs):

            x_arr = np.array([])
            y_arr = np.array([])

            if data is not None:
                if x:
                    x_arr = data[x].to_numpy()
                    y_arr = np.full((x_arr.size,), '', dtype=str)
                if y:
                    y_arr = data[y].to_numpy()
                    if x_arr.size == 0:
                        x_arr = np.full((y_arr.size,), '', dtype=str)

            drawing = DrawingBuilder.get_bar_drawing(
                x_arr, y_arr,
                DrawingLibrary.seaborn,
                kwargs,
            )
            drawings.append(drawing)

        def violinplot(*, x=None, y=None, data=None, **kwargs):

            if data is not None:
                if x is None and y is not None:
                    data = data[y]
                elif y is None and x is not None:
                    data = data[x]
                elif x is not None and y is not None:
                    data = pd.concat([data[x], data[y]], axis=1).reset_index()
            else:
                if x is None:
                    data = y
                elif y is None:
                    data = x
                else:
                    data = pd.concat([x, y], axis=1).reset_index()

            drawing = Drawing(
                DrawingLibrary.seaborn,
                DrawingType.violin,
                data,
                kwargs,
            )

            drawings.append(drawing)

        def heatmap(data=None, **kwargs):
            if data is None:
                return

            curr_data = {
                'x': np.array(data)
            }

            drawing = Drawing(
                DrawingLibrary.seaborn,
                DrawingType.heatmap,
                None,
                kwargs,
            )

            drawings.append(drawing)

        def boxplot(x=None, y=None, data=None, **kwargs):

            if data is None:
                curr_data = {
                    'x': np.array(x),
                    'y': np.array(y)
                }

                drawing = Drawing(
                    DrawingLibrary.seaborn,
                    DrawingType.box,
                    None,
                    kwargs,
                )

                drawings.append(drawing)
                return

            x_arr = []
            y_arr = []

            if not x and not y:
                for column in data.columns:
                    if not is_numeric_dtype(data[column]):
                        continue

                    curr_data = {
                        'x': np.array([column]),
                        'y': data[column].to_numpy()
                    }

                    drawing = Drawing(
                        DrawingLibrary.seaborn,
                        DrawingType.box,
                        None,
                        kwargs,
                    )
                    drawings.append(drawing)
                return

            if x:
                x_arr = data[x].to_numpy()
            if y:
                y_arr = data[y].to_numpy()

            curr_data = {
                'x': np.array(x_arr),
                'y': np.array(y_arr)
            }

            drawing = Drawing(
                DrawingLibrary.seaborn,
                DrawingType.box,
                None,
                kwargs,
            )
            drawings.append(drawing)

        if not SeabornHandler._saved:
            SeabornHandler._saved = True
            SeabornHandler._displot = sns.displot
            SeabornHandler._histplot = sns.histplot
            SeabornHandler._lineplot = sns.lineplot
            SeabornHandler._lmplot = sns.lmplot
            SeabornHandler._scatterplot = sns.scatterplot
            SeabornHandler._catplot = sns.catplot
            SeabornHandler._barplot = sns.barplot
            SeabornHandler._violinplot = sns.violinplot
            SeabornHandler._heatmap = sns.heatmap
            SeabornHandler._boxplot = sns.boxplot

        sns.displot = displot
        sns.histplot = histplot
        sns.lineplot = lineplot
        sns.lmplot = lmplot
        sns.scatterplot = scatterplot
        sns.catplot = catplot
        sns.barplot = barplot
        sns.violinplot = violinplot
        sns.heatmap = heatmap
        sns.boxplot = boxplot

        SeabornHandler._replaced = True

    @staticmethod
    def revert_plots():

        if not SeabornHandler._replaced:
            return

        MatplotlibHandler.revert_plots()

        import seaborn as sns

        sns.displot = SeabornHandler._displot
        sns.histplot = SeabornHandler._histplot
        sns.lineplot = SeabornHandler._lineplot
        sns.lmplot = SeabornHandler._lmplot
        sns.scatterplot = SeabornHandler._scatterplot
        sns.catplot = SeabornHandler._catplot
        sns.barplot = SeabornHandler._barplot
        sns.violinplot = SeabornHandler._violinplot
        sns.heatmap = SeabornHandler._heatmap
        sns.boxplot = SeabornHandler._boxplot

        reload(sns)

        SeabornHandler._replaced = False
