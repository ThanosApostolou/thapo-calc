project('thapo-calc', 'cpp')

gtkmm_dep = dependency('gtkmm-3.0', method : 'pkg-config')

lib_src = ['src/libthapo-calc/libthapo-calc.cpp', 'src/libthapo-calc/grammar.cpp']
incdir = include_directories('src/libthapo-calc')
thapo_calc_cli_src = ['src/thapo-calc-cli/thapo-calc-cli.cpp']
thapo_calc_gtkmm_src = ['src/thapo-calc-gtkmm/thapo-calc-gtkmm.cpp', 'src/thapo-calc-gtkmm/buttons.cpp']

lib = shared_library('thapo-calc', sources : lib_src, version : '0.2', soversion : '0.2', install : true)
executable('thapo-calc-cli', sources : thapo_calc_cli_src, include_directories : incdir, link_with : lib, install : true)
executable('thapo-calc-gtkmm', sources : thapo_calc_gtkmm_src, include_directories : incdir, link_with : lib, dependencies : gtkmm_dep, gui_app : true, install : true)

install_data('README.md', install_dir : join_paths(get_option('datadir'), 'thapo-calc'))
