<p>
	ColdFusion has the ability to communicate with a number of different databases, which will be covered later on in this course; for the sample application we will be working on throughout the course, we will be using MySQL. If you already have MySQL 4 or 5 already installed, you can proceed to the 'Install Sample Files' section. If not, follow the steps below:
</p>
<h2>
Windows
</h2>
<p>
	<ol>
		<li>
			Open up a browser and go to: <a href="http://dev.mysql.com/downloads/mysql/" target="_new">http://dev.mysql.com/downloads/mysql/</a>
		</li>
		<li>
			Scroll to the list of available downloads. Click the 'Download' button next to the applicable download.
<ul>
<li>
Windows 64 bit -  Windows (x86, 64-bit), MSI
</li>
<li>
Windows 32 bit - Windows (x86, 32-bit), MSI
</li>
</ul>
If you are unsure, assume you are on a 32 bit machine.
		</li>
		<li>
			To begin the download, you must either login using pre-existing credentials by clicking the 'Proceed' button under the New Users section or click the 'No thanks, just start my download' link at the bottom of the screen.
		</li>
		<li>
			Save the file to your desktop.
		</li>
		<li>
			Double click the file.
		</li>
		<li>
			On the Welcome screen click 'Next'.
		</li>
		<li>
			Accept the terms in the license agreement and click 'Next'.
		</li>
		<li>
			When presented with a list of available Setup Types, Select 'Typical'.
		</li>
		<li>
			When ready to install, click 'Install'.
		</li>
		<li>
			If a system message pops up asking if you want the software to be installed on your machine, click 'Yes'.
		</li>
		<li>
			If a MySQL enterprise pop up window appears, click 'Next' until it disappears.
		</li>
		<li>
			When the Setup Wizard has completed, Click 'Finish'.
		</li>
		<li>
			If you receive another pop up message asking if you want the software to be installed on your machine, click 'Yes'.
		</li>
		<li>
			When the Server Instance Configuration Wizard displays, Click 'Next'.
		</li>
		<li>
			When displayed the available Server Instance Configuration types select 'Standard Configuration' and click 'Next'.
		</li>
		<li>
			When on the Windows Options scree, keep 'Install as a Service' selected as well as 'Launch the MySQL Server automatically'. Select the 'Include Bin Directory in Windows PATH' option and select 'Next'.
		</li>
		<li>
			On the Security Options scree, specify a new root password and select 'Next'. Remember this password as it will be needed later!
		</li>
		<li>
			When the Ready to Execute screen displays, click 'Execute'.
		</li>
		<li>
			After all the Configuration steps have successfully run, Click 'Finish'.
		</li>
		<li>
			You have successfully installed MySQL.
		</li>
	</ol>
</p>
<p>
	<object width="640" height="360"><param name="movie" value="http://www.youtube.com/v/JGNP6PP2-j8?version=3&amp;hl=en_US"></param><param name="allowFullScreen" value="true"></param><param name="allowscriptaccess" value="always"></param><embed src="http://www.youtube.com/v/JGNP6PP2-j8?version=3&amp;hl=en_US" type="application/x-shockwave-flash" width="640" height="360" allowscriptaccess="always" allowfullscreen="true"></embed></object>
</p>
<h2>
	Mac
</h2>
<p>
	<ol>
		<li>
			Open up a browser and go to: <a href="http://dev.mysql.com/downloads/mysql/" target="_new">http://dev.mysql.com/downloads/mysql/</a>
		</li>
		<li>
			Scroll to the list of available downloads. Click the 'Download' button next to the applicable download.
			<ul>
				<li>
					Mac 64 bit - Mac OS X ver. 10.6 (x86, 64-bit), DMG Archive
				</li>
				<li>
					Mac 32 bit - Mac OS X ver. 10.6 (x86, 32-bit), DMG Archive
				</li>
			</ul>
			If you are unsure, assume you are on a 32 bit machine
		</li>
		<li>
			To begin the download, you must either login using pre-existing credentials by clicking the 'Proceed' button under the New Users section or click the 'No thanks, just start my download' link at the bottom of the screen.
		</li>
		<li>
			Save the file to your desktop.
		</li>
		<li>
			Double click the file.
		</li>
		<li>
			Once the DMG has mounted, double click the mysql-5.5.28-osx10.6-x86_64.pkg (or similar) file.  It should be the first file in the list.
		</li>
		<li>
			On the welcome screen, click 'Continue'.
		</li>
		<li>
			Review the Important Information screen and click 'Continue'.
		</li>
		<li>
			Review the Software License Agreement and click 'Continue'.
		</li>
		<li>
			Agree to the terms.
		</li>
		<li>
			Select the drive you wish to install the software to and click 'Continue'.
		</li>
		<li>
			Click 'Install'.
		</li>
		<li>
			Provide your system password and click 'OK'.
		</li>
		<li>
			Once you receive the screen stating the Installation was Successful, click 'Close'.
		</li>
		<li>
			Double click the MySQL.prefPane icon.
		</li>
		<li>
			Click 'Install'.
		</li>
		<li>
			If the MySQL server is not running, click the 'Start MySQL Server' button.
		</li>
		<li>
			If prompted, provide your system password and click 'OK'.
		</li>
		<li>
			Click the 'Automatically Start MySQL Server on Startup' check box.
		</li>
		<li>
			Close the window.
		</li>
		<li>
			Open up a terminal window and enter the following:<br />
<span class="code">/usr/local/mysql/bin/mysqladmin -u root password [NewPassword]</span>
		</li>
		<li>
			Make sure to replace [NewPassword] with the password you wish to use for the root user.  Remember this password, it will be needed later.
		</li>
		<li>
			Close the Terminal window.
		</li>
		<li>
			You have successfully installed MySQL.
		</li>
	</ol>
</p>
<p>
	<object width="640" height="360"><param name="movie" value="http://www.youtube.com/v/LVV78YWohA0?version=3&amp;hl=en_US"></param><param name="allowFullScreen" value="true"></param><param name="allowscriptaccess" value="always"></param><embed src="http://www.youtube.com/v/LVV78YWohA0?version=3&amp;hl=en_US" type="application/x-shockwave-flash" width="640" height="360" allowscriptaccess="always" allowfullscreen="true"></embed></object>
</p>
