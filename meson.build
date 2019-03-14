project('thapo-calc', 'cpp')

gtkmm_dep = dependency('gtkmm-3.0', method : 'pkg-config')

incdir = include_directories('src/include')
lib_src = ['src/thapo-calc-lib/thapo-calc-lib.cpp', 'src/thapo-calc-lib/grammar.cpp']
thapo_calc_cli_src = ['src/thapo-calc-cli/thapo-calc-cli.cpp']
thapo_calc_gtkmm_src = ['src/thapo-calc-gtkmm/thapo-calc-gtkmm.cpp', 'src/thapo-calc-gtkmm/buttons.cpp']

lib = shared_library('thapo-calc', sources : lib_src, include_directories : incdir, version : '0.2', soversion : '0.2', install : true)
executable('thapo-calc-cli', sources : thapo_calc_cli_src, include_directories : incdir, link_with : lib, install : true)
executable('thapo-calc-gtkmm', sources : thapo_calc_gtkmm_src, include_directories : incdir, link_with : lib, dependencies : gtkmm_dep, gui_app : true, install : true)

install_data('README.md')
if build_machine.system() == 'linux'
	install_data('data/thapo-calc-gtkmm.desktop', install_dir : join_paths(get_option('datadir'), 'applications'))
elif build_machine.system() == 'windows'
	meson.add_install_script('windows-bundle.sh', get_option('bindir'), get_option('libdir'))
endif
